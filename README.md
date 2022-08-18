# AutoGarf

This program performs copy/rename actions on files.

## Instructions:

### Step 1 - Install / Download Autogarf

Download the zip file from the [releases](https://github.com/matthewchivers/autogarf/releases) section of https://github.com/matthewchivers/autogarf.

Extract the contents to a meaningful location - such as `C:\Program Files\autogarf`

### Step 2 - Specify a directory

Let's say your directory structure looks something like this:

```
C:
---- user
    |---- documents
         |---- project
              |---- clients
                   |---- client A
                        |---- Statement 30 Apr 22.docx
                        |---- Statement 31 May 22.docx
                   |---- client B
                        |---- Statement 30 Apr 22.docx
                        |---- Statement 31 May 22.docx
                   |---- client C
                        |---- Statement 30 Apr 22.docx
                        |---- Statement 31 May 22.docx
```
Within the `clients` directory there area  series of directories (`client A`, `client B`, `client C`), each of these holds a Statement file.

The "top" directory in this instance is the "clients" directory, but it could be called something else.

In this example, the directory path will look like `C:\user\documents\project\clients`.

This is the directory that needs to be specified such that the program can recognise it and perform tasks on all the folders / files inside it.

#### Step 2.1 - Edit the config 
Edit the file `config.yaml` (which should already be in the same directory as `autogarf.exe`)

The file should have just one line when finished:

```
client-directory: 'C:\user\documents\project\clients'
```
> Make sure you use single quote to surround the directory path rather than backticks or double quotes.

### Step 3 - Run the program

Navigate to the directory that contains the executable `autogarf.exe` (this will be wherever you extracted it to in Step 1).

Double click on the executable - it should run.

You won't notice much happen - perhaps a black box will pop up, but it will be gone in the blink of an eye.
Most of the actions are designed to happen in the background.

How you'll know it has worked:

1) Files have been replecated and renamed in the directories you specified.
2) There will be a new directory (`logs`) at the same location where autogarf.exe is held.
   This new directory should hold a file (and eventually more than one file) which detail the exact steps the program has taken.
   The logs in here will be useful if anything goes wrong / doesn't work as intended.

### Step 4 - Enjoy

That's it! The program should work and all is well in the world

### Optional - Schedule to run automatically

1. Use the "Windows Key" + "R" to open run. Type "taskschd.msc" and press Enter on your keyboard. This will open Task Scheduler.
1. Under the actions panel, choose "Create Task".
1. The "Create Task" screen will appear. Select the "General" tab
1. In the "Name" field, give the task a name. Example: "AutoGarf".
1. In the "Description" field, here you can describe what the task is for and what it will do.
1. Select the "Triggers" tab.
1. Select "New…".
1. The "New Trigger" window will appear, here you have the option to set when the task will start.
1. Select when you would like the task to start in the "Begin the task" drop-down menu.
1. Modify the "Settings" area as desired.
1. "Enabled" (at the bottom) is checked by default (but please make sure)
1. Select "OK".
1. Select the "Actions" tab, then select "New".
1. The "New Action" window will open.
1. In the "Action" drop down, "Start a program" is set by default. Change it if desired.
1. Select "Browse…" next to the "Program/script" field
1. Navigate to (and select) `autogarf.exe`
1. In the box labelled "Start in (optional)" input the path to the directory where `autogarf.exe` is held.
    e.g. the same folder that is now in the Program/script box, but without `autogarf.exe` on the end.
1. Select "OK".
1. Go to the "Conditions" tab.
1. You can change these if you’d like, but I recommend leaving these settings default.
1. Select the "Settings" tab. 
1. Select "Run task as soon as possible after a scheduled start is missed"
1. Select "OK".



