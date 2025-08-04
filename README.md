# Vikunja-Github-Sync

This application is for syncing GitHub issues that are assigned to your user over to [Vikunja](vikunja.io).

It will create Vikunja Tasks for a GitHub issue under all the following conditions:
- You have set a default project in Vikunja
- The issue is assigned to your user
- The issue is open
- You have at least one comment in the issue or created the issue
- There was no previous task created for this issue

## Setup

You will need to set the following environment variables:
- **VIKUNJA_TOKEN** - Your api token for vikunja, needs the `Other:User`, `Tasks:Create` and `Tasks:Read One` permissions
- **VIKUNJA_URL** - The instance URL for vikunja

### Development

Feel free to create an `.env` file from the `example.env` preset and source it before executing.

## How it works

The application is designed to be as stateless as possible.
It should be possible run the application, create all the tasks that are not yet created in one go and end the application.
If there is the need to continuously create tasks the application can be used together with e.g. Systemd-Timers to run at a predefined scheduled.

The information if there is a task associated with for a certain GitHub issue is stored in the issue itself.
It resides in a hidden Markdown part in the last comment (may also be the first description of the issue) that your user has made in this issue at the state of execution of the application.
After fetching all the issues that you are assigned to, we check for each if there is already a corresponding task in Vikunja.
For the issues without corresponding task we create one.
