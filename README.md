# Set volume

Small console app that can set the current system volume.

I was tired of switching the volume between "music values" and "meeting values". The music values for me is low (10-12%) but for meetings is higher (70-90%).

Mostly used together with [PowerToys](https://learn.microsoft.com/en-us/windows/powertoys/).

## Set single value

`.\set-volume.exe 25` - set the current volume to 25%
`.\set-volume.exe 100` - set the current volume to 100% (max)

## Range of values

It is possible to pass range of values and the app will set the volume to the closest (higher) value based on the current system value. The range values are separated with `:`

Imagine that the current volume is `25`. Then the command below will set it to `60`. `60` is the next bigger number to `25`:

`.\set-volume 10:12:14:60:70:90`

Having the volume to `60` if we execute the same command the system volume will be set to `70`.
Running it again - `90`.
Running it again `10` - when the current volume is equal or greater than the max range value, from the list, the app will set the volume to the lowest from the list/range. This way we can loop through pre-defined volume levels with ease.

## PowerToys

The whole idea of this app was to be used with [PowerToys](https://learn.microsoft.com/en-us/windows/powertoys/). In there we can remap shortcut to run a program.

- start `PowerToys`
- choose `Keyboard Manager`
- scroll to `Remap a shortcut` and click on it
- press `+ Add shortcut remapping` (bottom left)
- choose the desired shortcut from the left side
- on the right side:

  - `Action` - `Run program`
  - `App` - the full path to `set-volume.exe`
  - `Args` - either single value (for example `20`) or range of values to loop. Separated by `:` (for example `10:12:14:60:70:100`)
  - `Start in` - (not sure its needed but ..) the folder where `set-volume.exe` is located
  - `Elevation` - `Normal`
  - `If running` - `Close`
  - `Visibility` - `Hidden`

    ![PowerToys config](./assets/powertoys_setup.png)

  - press `OK`

And thats it! (thinking about the screenshot above) if we now start pressing ``Ctrl + ` `` the volume will change between 10 -> 12 -> 14 -> 60 -> 70 -> 100 -> 10 ...
