# Veracode Slack Notifier Utility

## Description
Utility designed to be run in a build process after a Veracode scan to notify a Slack channel that the scan completed. Optionally, the notification can also include the compliance policy assigned to that app and whether or not it's passing.
For builds that don't wait for the Veracode scan to complete, the utility can be set to run on a schedule to provide notifications.

## Executables
Executables for Windows, Mac, and Linux will be available in the releases section of the repository (https://github.com/brian1917/vcodeSlackNotifier/releases)

## Running the Utility
The utility takes one argument - the location of the JSON config file. Run the utility as a command line
action at the end of the build (after Veracode completed):
`vCodeSlackNotifer appconfig.json`

## Configuration File
A sample config file is shown below. The 5 parameters below are require to be present.
```
{
    "credsFile": "/Users/bpitta/.veracode/credentials",
    "appID": "123456",
    "slackURL" : "https://hooks.slack.com/services/ABC12345/XYZ9876/abceEdfJhi",
    "onlyNotifyOnNotPass" : false,
    "includePolicyStatus" : true
}
```
In this example, the utility will provide a notifcation for App ID 123456 in the Veracode Platform for all builds (not just ones that don't pass), and will
include policy information in the notification.

## Integrating with Slack
Go to https://api.slack.com/ to create an app (I call it Veracode Notifer) and get your webhook URL. This is the URL you will use in the `slackURL` parameter of the config file.

## Third-party Packages
github.com/brian1917/vcodeapi (https://godoc.org/github.com/brian1917/vcodeapi)