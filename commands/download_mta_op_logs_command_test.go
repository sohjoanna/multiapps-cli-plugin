package commands_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	plugin_fakes "github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	baseclient "github.com/SAP/cf-mta-plugin/clients/baseclient"
	"github.com/SAP/cf-mta-plugin/clients/models"
	restfake "github.com/SAP/cf-mta-plugin/clients/restclient/fakes"
	slmpfake "github.com/SAP/cf-mta-plugin/clients/slmpclient/fakes"
	slppfake "github.com/SAP/cf-mta-plugin/clients/slppclient/fakes"
	"github.com/SAP/cf-mta-plugin/commands"
	cmd_fakes "github.com/SAP/cf-mta-plugin/commands/fakes"
	"github.com/SAP/cf-mta-plugin/testutil"
	"github.com/SAP/cf-mta-plugin/ui"
)

var _ = Describe("DownloadMtaOperationLogsCommand", func() {
	Describe("Execute", func() {
		const org = "test-org"
		const space = "test-space"
		const user = "test-user"

		var name string
		var cliConnection *plugin_fakes.FakeCliConnection
		var slmpClient *slmpfake.FakeSlmpClientOperations
		var slppClient *slppfake.FakeSlppClientOperations
		var restClient *restfake.FakeRestClientOperations
		var clientFactory *commands.TestClientFactory
		var command *commands.DownloadMtaOperationLogsCommand
		var oc = testutil.NewUIOutputCapturer()
		var ex = testutil.NewUIExpector()

		var getOutputLines = func(dir string) []string {
			wd, _ := os.Getwd()
			return []string{
				fmt.Sprintf("Downloading logs of multi-target app operation with id %s in org %s / space %s as %s...\n",
					testutil.ProcessID, org, space, user),
				"OK\n",
				fmt.Sprintf("Saving logs to %s"+string(os.PathSeparator)+"%s...\n", wd, dir),
				fmt.Sprintf("  %s\n", testutil.LogID),
				"OK\n",
			}
		}

		var expectDirWithLog = func(dir string) {
			Expect(exists(dir)).To(Equal(true))
			Expect(exists(dir + "/" + testutil.LogID)).To(Equal(true))
			Expect(contentOf(dir + "/" + testutil.LogID)).To(Equal(testutil.LogContent))
		}

		BeforeEach(func() {
			ui.DisableTerminalOutput(true)
			name = command.GetPluginCommand().Name
			cliConnection = cmd_fakes.NewFakeCliConnectionBuilder().
				CurrentOrg("test-org-guid", org, nil).
				CurrentSpace("test-space-guid", space, nil).
				Username(user, nil).
				AccessToken("bearer test-token", nil).
				APIEndpoint("https://api.test.ondemand.com", nil).Build()
			slmpClient = slmpfake.NewFakeSlmpClientBuilder().
				GetMetadata(&testutil.SlmpMetadataResult, nil).
				GetProcess(testutil.ProcessID, &testutil.ProcessResult, nil).Build()
			slppClient = slppfake.NewFakeSlppClientBuilder().
				GetMetadata(&testutil.SlppMetadataResult, nil).
				GetLogs(testutil.LogsResult, nil).
				GetLogContent(testutil.LogID, testutil.LogContent, nil).Build()
			restClient = restfake.NewFakeRestClientBuilder().Build()
			clientFactory = commands.NewTestClientFactory(slmpClient, slppClient, restClient)
			command = &commands.DownloadMtaOperationLogsCommand{}
			testTokenFactory := commands.NewTestTokenFactory(cliConnection)
			command.InitializeAll(name, cliConnection, testutil.NewCustomTransport(200, nil), nil, clientFactory, testTokenFactory)
		})

		// unknown flag - error
		Context("with an unknown flag", func() {
			It("should print incorrect usage, call cf help, and exit with a non-zero status", func() {
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"-a"}).ToInt()
				})
				ex.ExpectFailure(status, output, "Incorrect usage. Unknown or wrong flag.")
				Expect(cliConnection.CliCommandArgsForCall(0)).To(Equal([]string{"help", name}))
			})
		})

		// wrong arguments - error
		Context("with wrong arguments", func() {
			It("should print incorrect usage, call cf help, and exit with a non-zero status", func() {
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"x", "y", "z"}).ToInt()
				})
				ex.ExpectFailure(status, output, "Incorrect usage. Wrong arguments.")
				Expect(cliConnection.CliCommandArgsForCall(0)).To(Equal([]string{"help", name}))
			})
		})

		// no arguments - error
		Context("with no arguments", func() {
			It("should print incorrect usage, call cf help, and exit with a non-zero status", func() {
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{}).ToInt()
				})
				ex.ExpectFailure(status, output, "Incorrect usage. Missing required options '[i]'.")
				Expect(cliConnection.CliCommandArgsForCall(0)).To(Equal([]string{"help", name}))
			})
		})

		// can't connect to backend - error
		Context("when can't connect to backend", func() {
			const host = "x"
			It("should print an error and exit with a non-zero status", func() {
				clientFactory.SlmpClient = slmpfake.NewFakeSlmpClientBuilder().
					GetMetadata(nil, fmt.Errorf("Get https://%s/slprot/test/test/slp/metadata: dial tcp: lookup %s: no such host", host, host)).Build()
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"-i", testutil.ProcessID, "-u", host}).ToInt()
				})
				ex.ExpectFailureOnLine(status, output, "Could not get SLMP metadata:", 1)
			})
		})

		// backend returns an an error response (GetMetadata) - error
		Context("with an error response returned by the backend", func() {
			It("should print an error and exit with a non-zero status", func() {
				clientFactory.SlmpClient = slmpfake.NewFakeSlmpClientBuilder().
					GetMetadata(nil, fmt.Errorf("unknown error (status 404)")).Build()
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"-i", testutil.ProcessID}).ToInt()
				})
				ex.ExpectFailureOnLine(status, output, "Could not get SLMP metadata:", 1)
			})
		})

		// non-existing process id - error
		Context("with a non-existing process id", func() {
			It("should print an error and exit with a non-zero status", func() {
				var clientError = baseclient.NewClientError(testutil.ClientError)
				clientFactory.SlmpClient = slmpfake.NewFakeSlmpClientBuilder().
					GetProcess("test", nil, clientError).
					GetMetadata(&testutil.SlmpMetadataResult, nil).Build()
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"-i", "test"}).ToInt()
				})
				ex.ExpectFailureOnLine(status, output, "Multi-target app operation with id test not found", 1)
				Expect(exists("mta-op-test")).To(Equal(false))
			})
		})

		// existing process id, backend returns an error response (GetLogs) - error
		Context("with an existing process id and an error response returned by the backend", func() {
			It("should print an error and exit with a non-zero status", func() {
				clientFactory.SlppClient = slppfake.NewFakeSlppClientBuilder().
					GetMetadata(&testutil.SlppMetadataResult, nil).
					GetLogs(models.Logs{}, fmt.Errorf("unknown error (status 500)")).Build()
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"-i", testutil.ProcessID}).ToInt()
				})
				ex.ExpectFailureOnLine(status, output, "Could not get process logs:", 1)
				Expect(exists("mta-op-" + testutil.ProcessID)).To(Equal(false))
			})
		})

		// existing process id, backend returns an error response (GetLogContent) - error
		Context("with an existing process id and an error response returned by the backend", func() {
			It("should print an error and exit with a non-zero status", func() {
				fakeSlppClient := slppfake.NewFakeSlppClientBuilder().
					GetMetadata(&testutil.SlppMetadataResult, nil).
					GetLogs(testutil.LogsResult, nil).
					GetLogContent("", "", fmt.Errorf("unknown error (status 500)")).Build()
				clientFactory.SlppClient = fakeSlppClient
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"-i", testutil.ProcessID}).ToInt()
				})
				ex.ExpectFailureOnLine(status, output, fmt.Sprintf("Could not get content of log %s:", testutil.LogID), 1)
				Expect(fakeSlppClient.GetLogContentArgsForCall(0)).To(Equal(testutil.LogID))
				Expect(exists("mta-op-" + testutil.ProcessID)).To(Equal(false))
			})
		})

		// existing process id - success
		Context("with an existing process id", func() {
			const dir = "mta-op-" + testutil.ProcessID
			It("should download the logs for the current process and exit with zero status", func() {
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"-i", testutil.ProcessID}).ToInt()
				})
				ex.ExpectSuccessWithOutput(status, output, getOutputLines(dir))
				expectDirWithLog(dir)
			})
			AfterEach(func() {
				os.RemoveAll(dir)
			})
		})

		// existing process id and existing directory - error
		Context("with an existing process id and an existing directory", func() {
			const customDir string = "test"
			BeforeEach(func() {
				os.Mkdir(customDir, 0755)
			})
			It("should print an error and exit with a non-zero status", func() {
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"-i", testutil.ProcessID, "-d", customDir}).ToInt()
				})
				ex.ExpectFailureOnLine(status, output, fmt.Sprintf("Could not create download directory %s:", customDir), 2)
			})
			AfterEach(func() {
				os.RemoveAll(customDir)
			})
		})

		if runtime.GOOS != "windows" {
			// existing process id and a directory that can't be written to - error
			Context("with an existing process id and a directory that can't be written to", func() {
				const customDir string = "test"
				BeforeEach(func() {
					os.Mkdir(customDir, 0000)
				})
				It("should print an error and exit with a non-zero status", func() {
					output, status := oc.CaptureOutputAndStatus(func() int {
						return command.Execute([]string{"-i", testutil.ProcessID, "-d", customDir + "/subdir"}).ToInt()
					})
					ex.ExpectFailureOnLine(status, output, fmt.Sprintf("Could not save log %s:", testutil.LogID), 4)
				})
				AfterEach(func() {
					os.Chmod(customDir, 0755)
					os.RemoveAll(customDir)
				})
			})
		}

		// existing process id and non-existing directory - success
		Context("with an existing process id and a non-existing directory", func() {
			const customDir string = "test-non-existing"
			It("should create the directory, download the logs for the current process and exit with zero status", func() {
				output, status := oc.CaptureOutputAndStatus(func() int {
					return command.Execute([]string{"-i", testutil.ProcessID, "-d", customDir}).ToInt()
				})
				ex.ExpectSuccessWithOutput(status, output, getOutputLines(customDir))
				expectDirWithLog(customDir)
			})
			AfterEach(func() {
				os.RemoveAll(customDir)
			})
		})
	})
})

func contentOf(fileName string) string {
	content, _ := ioutil.ReadFile(fileName)
	return string(content)
}

func exists(dirName string) bool {
	_, err := os.Stat(dirName)
	if err == nil {
		return true
	}
	return false
}