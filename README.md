# lavie

**Load and Access a Vitualbox Image Easily** (___WIP___)

For development, I needed an access to a GNU/Linux server, I installed VirtualBox and Putty.

Each time, I have to launch a VM, wait, and then launch Putty (or access with ssh)

So I decided to make a program to do it automatically, I started with [Purebasic](http://www.purebasic.com),
great and nice language and made a small program to launch my VM automatically

<details><summary markdown="span"><code>LaunchVM.pb source code</code></summary>

```BlitzBasic
; For the directory get the env variable:
; VBOX_INSTALL_PATH or VBOX_MSI_INSTALL_PATH
;VBoxManage_dir.s = GetEnvironmentVariable("VBOX_INSTALL_PATH")
VBoxManage_dir.s = GetEnvironmentVariable("VBOX_MSI_INSTALL_PATH")
;VBoxManage_dir.s = "c:\Program Files\Oracle\VirtualBox\"

VBoxManage_program.s = "VBoxManage.exe"

vm_name.s = Chr(34) + "Ubuntu 16.04.3 Server" + Chr(34)
vm_username.s = "your_username"
vm_password.s = "your_password"

; "c:\Program Files\Oracle\VirtualBox\VBoxManage.exe" startvm "Ubuntu 16.04.3 Server" --type headless
key_start.s = "startvm"
key_start_headless.s = "--type headless"

; "c:\Program Files\Oracle\VirtualBox\VBoxManage.exe" guestproperty get "Ubuntu 16.04.3 Server" /VirtualBox/GuestInfo/Net/0/V4/IP
key_search.s = "guestproperty get"
key_search_IP.s = "/VirtualBox/GuestInfo/Net/0/V4/IP"

; "c:\Program Files\Oracle\VirtualBox\VBoxManage.exe" controlvm "Ubuntu 16.04.3 Server" poweroff
key_stop.s = "controlvm"
key_stop_poweroff.s = "poweroff" ;"acpipowerbutton"

key_user.s = "--username " + vm_username + " --password " + vm_password

parameter_start.s = key_start + " " + vm_name + " " + key_start_headless
parameter_check.s = key_search + " " + vm_name + " " + key_search_IP
parameter_stop.s = key_stop + " " + vm_name + " " + key_stop_poweroff

; C:\app\putty\PUTTY.EXE -load "Ubuntu 16.04.3 LTS" -l <vm_username>
putty_program_dir.s = "C:\app\tools\putty\"
putty_program.s = "PUTTY.EXE"
putty_param.s = "-load"

putty_session.s = Chr(34) + "Ubuntu 16.04.3 LTS" + Chr(34)

putty_vm_user_login.s = "-l " + vm_username + " -pw " + vm_password

putty_parameter.s = putty_param + " " + putty_session + " " + putty_vm_user_login

get_value.i = #False
output.s = ""

Procedure.s _readProgramData(ProgramID.i)
    Protected result.s

    *buffer = AllocateMemory(250)
    ReadProgramData(ProgramID, *buffer, 250)
    result = PeekS(*buffer, -1, #PB_Ascii)
    FreeMemory(*buffer)

    ProcedureReturn result
EndProcedure

If OpenConsole("LaunchVM")
    PrintN("LaunchVM v0.16"):PrintN("")

    ; First run VM
    PrintN("Waiting for VM " + vm_name + " to power on...")
    VBoxManage_start.i = RunProgram(VBoxManage_dir + VBoxManage_program, parameter_start, "", #PB_Program_Hide | #PB_Program_Wait)

    If VBoxManage_start
        ; Checking
        Print("Waiting for the VM to be ready to connect")
        VBoxManage_check.i = RunProgram(VBoxManage_dir + VBoxManage_program, parameter_check, "", #PB_Program_Open | #PB_Program_Read | #PB_Program_Error)
        If VBoxManage_check
            If AvailableProgramOutput(VBoxManage_check)
                output = ReadProgramString(VBoxManage_check)
            Else
                output = _readProgramData(VBoxManage_check)
            EndIf
            output = ReplaceString(output, Chr(13)+Chr(10), "")
            CloseProgram(VBoxManage_check)

            While get_value = #False
                If output <> "No value set!"
                    get_value = #True
                Else
                    VBoxManage_check.i = RunProgram(VBoxManage_dir + VBoxManage_program, parameter_check, "", #PB_Program_Open | #PB_Program_Read | #PB_Program_Error)
                    If AvailableProgramOutput(VBoxManage_check)
                        output = ReadProgramString(VBoxManage_check)
                    Else
                        output = _readProgramData(VBoxManage_check)
                    EndIf
                    output = ReplaceString(output, Chr(13)+Chr(10), "")
                    CloseProgram(VBoxManage_check)

                    If Random(4) = 1 : Print(".") : EndIf
                    ; Just wait a little before checking again
                    Delay(250)
                EndIf
            Wend

            IP.s = Mid(output, 8)
            PrintN("")
            PrintN("IP " + IP + " is UP, ready to connect...")

            PrintN("Launching PUTTY with session " + putty_session)
            putty_prg = RunProgram(putty_program_dir + putty_program, putty_parameter, "", #PB_Program_Wait)
            If putty_prg = 0
                PrintN("ERROR: Cannot launch " + Chr(34) + putty_program_dir + putty_program + Chr(34))
            EndIf

            ; Because we're using #PB_Program_Wait for putty, in the console it's waiting
            ; We quitted PUTTY, we can stop the vm
            PrintN("Waiting for VM " + vm_name + " to power off...")
            RunProgram(VBoxManage_dir + VBoxManage_program, parameter_stop, "", #PB_Program_Hide | #PB_Program_Wait)
            PrintN("...Done")
        EndIf
    EndIf
    CloseConsole()
EndIf
```

</details>

- - - -

As it only launches ___ONE___ VM with my session I created on Putty, I was thinking to use the registry, json, yaml, ...

Purebasic can handle it, but my code started to be messy :stuck_out_tongue_winking_eye:

So here I am now, and I choose [Golang](https://golang.org/), because I'm starting to learn it and I like it :smile:

- - - -

## __TO DO__

- [ ] Create a VM (will become ___clavie___)
- - - -
