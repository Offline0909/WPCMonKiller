# WPCMonKiller
Discreet program for disabling Window's 10 parental control features without any heavy-handed approaches.

# What does it do

Constantly checks for if any Windows Parental Control Monitor processes have started and instantly terminates them.
Has worked for years; however, Windows recently has picked up on this and includes a warning about Administrator priveleges on a child user-account.

Will completely negate:
* Application screen time (how long can be spent on an app per set quota)
* Application blocks (ability/inability to use certain applications)
* Device screen time (how long can be spent on computer etc.)
* Application logging (weekly parent report will show as 0 time spent)

# Requirements
Will require Administrator privileges on the user account. Without them, you will not be able to kill the WPCMon task.
For more optimal use, give program administrative priveleges and place in startup folder.
