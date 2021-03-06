# Getting started

Once you have [installed](../README.md#install) azbrowse you should simply be able to run `azbrowse` from your terminal/command line.

## Layout and navigation

Assuming you haven't yet customised your key bindings, you can use the cursor keys to move the selected item, Enter to select and backspace go back.

The interface for azbrowse has three main sections illustrated below. On the left is the list view - the up and down arrows let you move up and down the list, and you can select items with `Enter`. When you select an item it is shown in the content panel on the right.

At the bottom of the screen is the status bar where messages are displayed to show the progress of actions and errors.

![layout](images/layout.jpg)

Note that there are a few shortcut keys displayed as prompts in the UI, e.g. `Ctrl+I` for help.

![help](images/help.jpg)

## Commands

`Ctrl+P` will bring up the command palette which shows you what options you have for the currently expanded item as shown below. Where commands have a key binding that is show at the right.

![command palette](images/command-palette.jpg)

## Guided tours

This section has walk-through videos to guide you through some aspects of azbrowse.

### Filtering, commands and `--navigate`

The up and down arrows are good for exploring the resources that you have, but as you become familiar with where you want to navigate to the filter command can be really helpful. To start, type `/` to bring up the filter panel and then start typing - the items in the list view will be filtered as you type. You can then press `Enter` or `Esc` to dismiss the filter panel and then either navigate the list with cursor keys and `Enter` or clear the filter with `Esc`.

Another great navigation feature is the `--navigate` command line argument which lets you auto-navigate to a resource. To get the ID for a resource you can use the "Copy current resource ID" command from the command palette (`Ctrl+P`)

![filtering, commands and --navigate](images/azbrowse-navigate-copyid-filter.gif)

### Updating content

For resources that have `PUT` endpoints defined in their API specs, azbrowse allows you to edit the content and send the update.

For example, you can navigate to a site in Azure App Service and then drill in to `config/appsettings` to see the current settings for the site. `Ctrl+U` can then be used to open your configured editor (by default it tries to use Visual Studio code but it is [configurable](./config.md#editing-content)). When you save and close the file, azbrowse will issue the `PUT` request with the new content. If you don't want to make a change then you can close the file without changes, or delete the file content and azbrowse will skip applying the change.

![updating content](images/azbrowse-update.gif)

### Metrics

Lots of resources in Azure have metrics defined for them, and azbrowse has support for charting single-value metrics. Simple navigate to the `[Metrics]` node for a resource and pick a metric to display.

![displaying metrics](images/azbrowse-metrics.gif)
