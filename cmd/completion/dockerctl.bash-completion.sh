# bash completion for dockerctl                            -*- shell-script -*-

__dockerctl_debug()
{
    if [[ -n ${BASH_COMP_DEBUG_FILE} ]]; then
        echo "$*" >> "${BASH_COMP_DEBUG_FILE}"
    fi
}

# Homebrew on Macs have version 1.3 of bash-completion which doesn't include
# _init_completion. This is a very minimal version of that function.
__dockerctl_init_completion()
{
    COMPREPLY=()
    _get_comp_words_by_ref "$@" cur prev words cword
}

__dockerctl_index_of_word()
{
    local w word=$1
    shift
    index=0
    for w in "$@"; do
        [[ $w = "$word" ]] && return
        index=$((index+1))
    done
    index=-1
}

__dockerctl_contains_word()
{
    local w word=$1; shift
    for w in "$@"; do
        [[ $w = "$word" ]] && return
    done
    return 1
}

__dockerctl_handle_go_custom_completion()
{
    __dockerctl_debug "${FUNCNAME[0]}: cur is ${cur}, words[*] is ${words[*]}, #words[@] is ${#words[@]}"

    local shellCompDirectiveError=1
    local shellCompDirectiveNoSpace=2
    local shellCompDirectiveNoFileComp=4
    local shellCompDirectiveFilterFileExt=8
    local shellCompDirectiveFilterDirs=16

    local out requestComp lastParam lastChar comp directive args

    # Prepare the command to request completions for the program.
    # Calling ${words[0]} instead of directly dockerctl allows to handle aliases
    args=("${words[@]:1}")
    requestComp="${words[0]} __completeNoDesc ${args[*]}"

    lastParam=${words[$((${#words[@]}-1))]}
    lastChar=${lastParam:$((${#lastParam}-1)):1}
    __dockerctl_debug "${FUNCNAME[0]}: lastParam ${lastParam}, lastChar ${lastChar}"

    if [ -z "${cur}" ] && [ "${lastChar}" != "=" ]; then
        # If the last parameter is complete (there is a space following it)
        # We add an extra empty parameter so we can indicate this to the go method.
        __dockerctl_debug "${FUNCNAME[0]}: Adding extra empty parameter"
        requestComp="${requestComp} \"\""
    fi

    __dockerctl_debug "${FUNCNAME[0]}: calling ${requestComp}"
    # Use eval to handle any environment variables and such
    out=$(eval "${requestComp}" 2>/dev/null)

    # Extract the directive integer at the very end of the output following a colon (:)
    directive=${out##*:}
    # Remove the directive
    out=${out%:*}
    if [ "${directive}" = "${out}" ]; then
        # There is not directive specified
        directive=0
    fi
    __dockerctl_debug "${FUNCNAME[0]}: the completion directive is: ${directive}"
    __dockerctl_debug "${FUNCNAME[0]}: the completions are: ${out[*]}"

    if [ $((directive & shellCompDirectiveError)) -ne 0 ]; then
        # Error code.  No completion.
        __dockerctl_debug "${FUNCNAME[0]}: received error from custom completion go code"
        return
    else
        if [ $((directive & shellCompDirectiveNoSpace)) -ne 0 ]; then
            if [[ $(type -t compopt) = "builtin" ]]; then
                __dockerctl_debug "${FUNCNAME[0]}: activating no space"
                compopt -o nospace
            fi
        fi
        if [ $((directive & shellCompDirectiveNoFileComp)) -ne 0 ]; then
            if [[ $(type -t compopt) = "builtin" ]]; then
                __dockerctl_debug "${FUNCNAME[0]}: activating no file completion"
                compopt +o default
            fi
        fi
    fi

    if [ $((directive & shellCompDirectiveFilterFileExt)) -ne 0 ]; then
        # File extension filtering
        local fullFilter filter filteringCmd
        # Do not use quotes around the $out variable or else newline
        # characters will be kept.
        for filter in ${out[*]}; do
            fullFilter+="$filter|"
        done

        filteringCmd="_filedir $fullFilter"
        __dockerctl_debug "File filtering command: $filteringCmd"
        $filteringCmd
    elif [ $((directive & shellCompDirectiveFilterDirs)) -ne 0 ]; then
        # File completion for directories only
        local subDir
        # Use printf to strip any trailing newline
        subdir=$(printf "%s" "${out[0]}")
        if [ -n "$subdir" ]; then
            __dockerctl_debug "Listing directories in $subdir"
            __dockerctl_handle_subdirs_in_dir_flag "$subdir"
        else
            __dockerctl_debug "Listing directories in ."
            _filedir -d
        fi
    else
        while IFS='' read -r comp; do
            COMPREPLY+=("$comp")
        done < <(compgen -W "${out[*]}" -- "$cur")
    fi
}

__dockerctl_handle_reply()
{
    __dockerctl_debug "${FUNCNAME[0]}"
    local comp
    case $cur in
        -*)
            if [[ $(type -t compopt) = "builtin" ]]; then
                compopt -o nospace
            fi
            local allflags
            if [ ${#must_have_one_flag[@]} -ne 0 ]; then
                allflags=("${must_have_one_flag[@]}")
            else
                allflags=("${flags[*]} ${two_word_flags[*]}")
            fi
            while IFS='' read -r comp; do
                COMPREPLY+=("$comp")
            done < <(compgen -W "${allflags[*]}" -- "$cur")
            if [[ $(type -t compopt) = "builtin" ]]; then
                [[ "${COMPREPLY[0]}" == *= ]] || compopt +o nospace
            fi

            # complete after --flag=abc
            if [[ $cur == *=* ]]; then
                if [[ $(type -t compopt) = "builtin" ]]; then
                    compopt +o nospace
                fi

                local index flag
                flag="${cur%=*}"
                __dockerctl_index_of_word "${flag}" "${flags_with_completion[@]}"
                COMPREPLY=()
                if [[ ${index} -ge 0 ]]; then
                    PREFIX=""
                    cur="${cur#*=}"
                    ${flags_completion[${index}]}
                    if [ -n "${ZSH_VERSION}" ]; then
                        # zsh completion needs --flag= prefix
                        eval "COMPREPLY=( \"\${COMPREPLY[@]/#/${flag}=}\" )"
                    fi
                fi
            fi
            return 0;
            ;;
    esac

    # check if we are handling a flag with special work handling
    local index
    __dockerctl_index_of_word "${prev}" "${flags_with_completion[@]}"
    if [[ ${index} -ge 0 ]]; then
        ${flags_completion[${index}]}
        return
    fi

    # we are parsing a flag and don't have a special handler, no completion
    if [[ ${cur} != "${words[cword]}" ]]; then
        return
    fi

    local completions
    completions=("${commands[@]}")
    if [[ ${#must_have_one_noun[@]} -ne 0 ]]; then
        completions+=("${must_have_one_noun[@]}")
    elif [[ -n "${has_completion_function}" ]]; then
        # if a go completion function is provided, defer to that function
        __dockerctl_handle_go_custom_completion
    fi
    if [[ ${#must_have_one_flag[@]} -ne 0 ]]; then
        completions+=("${must_have_one_flag[@]}")
    fi
    while IFS='' read -r comp; do
        COMPREPLY+=("$comp")
    done < <(compgen -W "${completions[*]}" -- "$cur")

    if [[ ${#COMPREPLY[@]} -eq 0 && ${#noun_aliases[@]} -gt 0 && ${#must_have_one_noun[@]} -ne 0 ]]; then
        while IFS='' read -r comp; do
            COMPREPLY+=("$comp")
        done < <(compgen -W "${noun_aliases[*]}" -- "$cur")
    fi

    if [[ ${#COMPREPLY[@]} -eq 0 ]]; then
		if declare -F __dockerctl_custom_func >/dev/null; then
			# try command name qualified custom func
			__dockerctl_custom_func
		else
			# otherwise fall back to unqualified for compatibility
			declare -F __custom_func >/dev/null && __custom_func
		fi
    fi

    # available in bash-completion >= 2, not always present on macOS
    if declare -F __ltrim_colon_completions >/dev/null; then
        __ltrim_colon_completions "$cur"
    fi

    # If there is only 1 completion and it is a flag with an = it will be completed
    # but we don't want a space after the =
    if [[ "${#COMPREPLY[@]}" -eq "1" ]] && [[ $(type -t compopt) = "builtin" ]] && [[ "${COMPREPLY[0]}" == --*= ]]; then
       compopt -o nospace
    fi
}

# The arguments should be in the form "ext1|ext2|extn"
__dockerctl_handle_filename_extension_flag()
{
    local ext="$1"
    _filedir "@(${ext})"
}

__dockerctl_handle_subdirs_in_dir_flag()
{
    local dir="$1"
    pushd "${dir}" >/dev/null 2>&1 && _filedir -d && popd >/dev/null 2>&1 || return
}

__dockerctl_handle_flag()
{
    __dockerctl_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    # if a command required a flag, and we found it, unset must_have_one_flag()
    local flagname=${words[c]}
    local flagvalue
    # if the word contained an =
    if [[ ${words[c]} == *"="* ]]; then
        flagvalue=${flagname#*=} # take in as flagvalue after the =
        flagname=${flagname%=*} # strip everything after the =
        flagname="${flagname}=" # but put the = back
    fi
    __dockerctl_debug "${FUNCNAME[0]}: looking for ${flagname}"
    if __dockerctl_contains_word "${flagname}" "${must_have_one_flag[@]}"; then
        must_have_one_flag=()
    fi

    # if you set a flag which only applies to this command, don't show subcommands
    if __dockerctl_contains_word "${flagname}" "${local_nonpersistent_flags[@]}"; then
      commands=()
    fi

    # keep flag value with flagname as flaghash
    # flaghash variable is an associative array which is only supported in bash > 3.
    if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
        if [ -n "${flagvalue}" ] ; then
            flaghash[${flagname}]=${flagvalue}
        elif [ -n "${words[ $((c+1)) ]}" ] ; then
            flaghash[${flagname}]=${words[ $((c+1)) ]}
        else
            flaghash[${flagname}]="true" # pad "true" for bool flag
        fi
    fi

    # skip the argument to a two word flag
    if [[ ${words[c]} != *"="* ]] && __dockerctl_contains_word "${words[c]}" "${two_word_flags[@]}"; then
			  __dockerctl_debug "${FUNCNAME[0]}: found a flag ${words[c]}, skip the next argument"
        c=$((c+1))
        # if we are looking for a flags value, don't show commands
        if [[ $c -eq $cword ]]; then
            commands=()
        fi
    fi

    c=$((c+1))

}

__dockerctl_handle_noun()
{
    __dockerctl_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    if __dockerctl_contains_word "${words[c]}" "${must_have_one_noun[@]}"; then
        must_have_one_noun=()
    elif __dockerctl_contains_word "${words[c]}" "${noun_aliases[@]}"; then
        must_have_one_noun=()
    fi

    nouns+=("${words[c]}")
    c=$((c+1))
}

__dockerctl_handle_command()
{
    __dockerctl_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"

    local next_command
    if [[ -n ${last_command} ]]; then
        next_command="_${last_command}_${words[c]//:/__}"
    else
        if [[ $c -eq 0 ]]; then
            next_command="_dockerctl_root_command"
        else
            next_command="_${words[c]//:/__}"
        fi
    fi
    c=$((c+1))
    __dockerctl_debug "${FUNCNAME[0]}: looking for ${next_command}"
    declare -F "$next_command" >/dev/null && $next_command
}

__dockerctl_handle_word()
{
    if [[ $c -ge $cword ]]; then
        __dockerctl_handle_reply
        return
    fi
    __dockerctl_debug "${FUNCNAME[0]}: c is $c words[c] is ${words[c]}"
    if [[ "${words[c]}" == -* ]]; then
        __dockerctl_handle_flag
    elif __dockerctl_contains_word "${words[c]}" "${commands[@]}"; then
        __dockerctl_handle_command
    elif [[ $c -eq 0 ]]; then
        __dockerctl_handle_command
    elif __dockerctl_contains_word "${words[c]}" "${command_aliases[@]}"; then
        # aliashash variable is an associative array which is only supported in bash > 3.
        if [[ -z "${BASH_VERSION}" || "${BASH_VERSINFO[0]}" -gt 3 ]]; then
            words[c]=${aliashash[${words[c]}]}
            __dockerctl_handle_command
        else
            __dockerctl_handle_noun
        fi
    else
        __dockerctl_handle_noun
    fi
    __dockerctl_handle_word
}

_dockerctl_completion()
{
    last_command="dockerctl_completion"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--help")
    flags+=("-h")
    local_nonpersistent_flags+=("--help")
    local_nonpersistent_flags+=("-h")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    must_have_one_noun+=("bash")
    must_have_one_noun+=("fish")
    must_have_one_noun+=("powershell")
    must_have_one_noun+=("zsh")
    noun_aliases=()
}

_dockerctl_config_ConfigCreate()
{
    last_command="dockerctl_config_ConfigCreate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--body=")
    two_word_flags+=("--body")
    flags+=("--configCreateBody.Data=")
    two_word_flags+=("--configCreateBody.Data")
    flags+=("--configCreateBody.Name=")
    two_word_flags+=("--configCreateBody.Name")
    flags+=("--configCreateBody.Templating.Name=")
    two_word_flags+=("--configCreateBody.Templating.Name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_config_ConfigDelete()
{
    last_command="dockerctl_config_ConfigDelete"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_config_ConfigInspect()
{
    last_command="dockerctl_config_ConfigInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_config_ConfigList()
{
    last_command="dockerctl_config_ConfigList"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_config_ConfigUpdate()
{
    last_command="dockerctl_config_ConfigUpdate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--body=")
    two_word_flags+=("--body")
    flags+=("--configSpec.Data=")
    two_word_flags+=("--configSpec.Data")
    flags+=("--configSpec.Name=")
    two_word_flags+=("--configSpec.Name")
    flags+=("--configSpec.Templating.Name=")
    two_word_flags+=("--configSpec.Templating.Name")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--version=")
    two_word_flags+=("--version")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_config()
{
    last_command="dockerctl_config"

    command_aliases=()

    commands=()
    commands+=("ConfigCreate")
    commands+=("ConfigDelete")
    commands+=("ConfigInspect")
    commands+=("ConfigList")
    commands+=("ConfigUpdate")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerArchive()
{
    last_command="dockerctl_container_ContainerArchive"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--path=")
    two_word_flags+=("--path")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerArchiveInfo()
{
    last_command="dockerctl_container_ContainerArchiveInfo"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--path=")
    two_word_flags+=("--path")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerAttach()
{
    last_command="dockerctl_container_ContainerAttach"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--detachKeys=")
    two_word_flags+=("--detachKeys")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--logs")
    flags+=("--stderr")
    flags+=("--stdin")
    flags+=("--stdout")
    flags+=("--stream")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerAttachWebsocket()
{
    last_command="dockerctl_container_ContainerAttachWebsocket"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--detachKeys=")
    two_word_flags+=("--detachKeys")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--logs")
    flags+=("--stderr")
    flags+=("--stdin")
    flags+=("--stdout")
    flags+=("--stream")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerChanges()
{
    last_command="dockerctl_container_ContainerChanges"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerCreate()
{
    last_command="dockerctl_container_ContainerCreate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--body=")
    two_word_flags+=("--body")
    flags+=("--containerCreateBody.ArgsEscaped")
    flags+=("--containerCreateBody.AttachStderr")
    flags+=("--containerCreateBody.AttachStdin")
    flags+=("--containerCreateBody.AttachStdout")
    flags+=("--containerCreateBody.Domainname=")
    two_word_flags+=("--containerCreateBody.Domainname")
    flags+=("--containerCreateBody.Healthcheck.Interval=")
    two_word_flags+=("--containerCreateBody.Healthcheck.Interval")
    flags+=("--containerCreateBody.Healthcheck.Retries=")
    two_word_flags+=("--containerCreateBody.Healthcheck.Retries")
    flags+=("--containerCreateBody.Healthcheck.StartPeriod=")
    two_word_flags+=("--containerCreateBody.Healthcheck.StartPeriod")
    flags+=("--containerCreateBody.Healthcheck.Timeout=")
    two_word_flags+=("--containerCreateBody.Healthcheck.Timeout")
    flags+=("--containerCreateBody.Hostname=")
    two_word_flags+=("--containerCreateBody.Hostname")
    flags+=("--containerCreateBody.Image=")
    two_word_flags+=("--containerCreateBody.Image")
    flags+=("--containerCreateBody.MacAddress=")
    two_word_flags+=("--containerCreateBody.MacAddress")
    flags+=("--containerCreateBody.NetworkDisabled")
    flags+=("--containerCreateBody.OpenStdin")
    flags+=("--containerCreateBody.StdinOnce")
    flags+=("--containerCreateBody.StopSignal=")
    two_word_flags+=("--containerCreateBody.StopSignal")
    flags+=("--containerCreateBody.StopTimeout=")
    two_word_flags+=("--containerCreateBody.StopTimeout")
    flags+=("--containerCreateBody.Tty")
    flags+=("--containerCreateBody.User=")
    two_word_flags+=("--containerCreateBody.User")
    flags+=("--containerCreateBody.WorkingDir=")
    two_word_flags+=("--containerCreateBody.WorkingDir")
    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerDelete()
{
    last_command="dockerctl_container_ContainerDelete"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--force")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--link")
    flags+=("--v")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerExport()
{
    last_command="dockerctl_container_ContainerExport"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerInspect()
{
    last_command="dockerctl_container_ContainerInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--size")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerKill()
{
    last_command="dockerctl_container_ContainerKill"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--signal=")
    two_word_flags+=("--signal")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerList()
{
    last_command="dockerctl_container_ContainerList"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--all")
    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--limit=")
    two_word_flags+=("--limit")
    flags+=("--size")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerLogs()
{
    last_command="dockerctl_container_ContainerLogs"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--follow")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--since=")
    two_word_flags+=("--since")
    flags+=("--stderr")
    flags+=("--stdout")
    flags+=("--tail=")
    two_word_flags+=("--tail")
    flags+=("--timestamps")
    flags+=("--until=")
    two_word_flags+=("--until")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerPause()
{
    last_command="dockerctl_container_ContainerPause"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerPrune()
{
    last_command="dockerctl_container_ContainerPrune"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerRename()
{
    last_command="dockerctl_container_ContainerRename"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerResize()
{
    last_command="dockerctl_container_ContainerResize"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--h=")
    two_word_flags+=("--h")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--w=")
    two_word_flags+=("--w")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerRestart()
{
    last_command="dockerctl_container_ContainerRestart"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--t=")
    two_word_flags+=("--t")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerStart()
{
    last_command="dockerctl_container_ContainerStart"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--detachKeys=")
    two_word_flags+=("--detachKeys")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerStats()
{
    last_command="dockerctl_container_ContainerStats"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--stream")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerStop()
{
    last_command="dockerctl_container_ContainerStop"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--t=")
    two_word_flags+=("--t")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerTop()
{
    last_command="dockerctl_container_ContainerTop"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--ps_args=")
    two_word_flags+=("--ps_args")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerUnpause()
{
    last_command="dockerctl_container_ContainerUnpause"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerUpdate()
{
    last_command="dockerctl_container_ContainerUpdate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--containerUpdateBody.BlkioWeight=")
    two_word_flags+=("--containerUpdateBody.BlkioWeight")
    flags+=("--containerUpdateBody.CgroupParent=")
    two_word_flags+=("--containerUpdateBody.CgroupParent")
    flags+=("--containerUpdateBody.CpuCount=")
    two_word_flags+=("--containerUpdateBody.CpuCount")
    flags+=("--containerUpdateBody.CpuPercent=")
    two_word_flags+=("--containerUpdateBody.CpuPercent")
    flags+=("--containerUpdateBody.CpuPeriod=")
    two_word_flags+=("--containerUpdateBody.CpuPeriod")
    flags+=("--containerUpdateBody.CpuQuota=")
    two_word_flags+=("--containerUpdateBody.CpuQuota")
    flags+=("--containerUpdateBody.CpuRealtimePeriod=")
    two_word_flags+=("--containerUpdateBody.CpuRealtimePeriod")
    flags+=("--containerUpdateBody.CpuRealtimeRuntime=")
    two_word_flags+=("--containerUpdateBody.CpuRealtimeRuntime")
    flags+=("--containerUpdateBody.CpuShares=")
    two_word_flags+=("--containerUpdateBody.CpuShares")
    flags+=("--containerUpdateBody.CpusetCpus=")
    two_word_flags+=("--containerUpdateBody.CpusetCpus")
    flags+=("--containerUpdateBody.CpusetMems=")
    two_word_flags+=("--containerUpdateBody.CpusetMems")
    flags+=("--containerUpdateBody.IOMaximumBandwidth=")
    two_word_flags+=("--containerUpdateBody.IOMaximumBandwidth")
    flags+=("--containerUpdateBody.IOMaximumIOps=")
    two_word_flags+=("--containerUpdateBody.IOMaximumIOps")
    flags+=("--containerUpdateBody.Init")
    flags+=("--containerUpdateBody.KernelMemory=")
    two_word_flags+=("--containerUpdateBody.KernelMemory")
    flags+=("--containerUpdateBody.KernelMemoryTCP=")
    two_word_flags+=("--containerUpdateBody.KernelMemoryTCP")
    flags+=("--containerUpdateBody.Memory=")
    two_word_flags+=("--containerUpdateBody.Memory")
    flags+=("--containerUpdateBody.MemoryReservation=")
    two_word_flags+=("--containerUpdateBody.MemoryReservation")
    flags+=("--containerUpdateBody.MemorySwap=")
    two_word_flags+=("--containerUpdateBody.MemorySwap")
    flags+=("--containerUpdateBody.MemorySwappiness=")
    two_word_flags+=("--containerUpdateBody.MemorySwappiness")
    flags+=("--containerUpdateBody.NanoCPUs=")
    two_word_flags+=("--containerUpdateBody.NanoCPUs")
    flags+=("--containerUpdateBody.OomKillDisable")
    flags+=("--containerUpdateBody.PidsLimit=")
    two_word_flags+=("--containerUpdateBody.PidsLimit")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--update=")
    two_word_flags+=("--update")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_ContainerWait()
{
    last_command="dockerctl_container_ContainerWait"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--condition=")
    two_word_flags+=("--condition")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container_PutContainerArchive()
{
    last_command="dockerctl_container_PutContainerArchive"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--copyUIDGID=")
    two_word_flags+=("--copyUIDGID")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--noOverwriteDirNonDir=")
    two_word_flags+=("--noOverwriteDirNonDir")
    flags+=("--path=")
    two_word_flags+=("--path")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_container()
{
    last_command="dockerctl_container"

    command_aliases=()

    commands=()
    commands+=("ContainerArchive")
    commands+=("ContainerArchiveInfo")
    commands+=("ContainerAttach")
    commands+=("ContainerAttachWebsocket")
    commands+=("ContainerChanges")
    commands+=("ContainerCreate")
    commands+=("ContainerDelete")
    commands+=("ContainerExport")
    commands+=("ContainerInspect")
    commands+=("ContainerKill")
    commands+=("ContainerList")
    commands+=("ContainerLogs")
    commands+=("ContainerPause")
    commands+=("ContainerPrune")
    commands+=("ContainerRename")
    commands+=("ContainerResize")
    commands+=("ContainerRestart")
    commands+=("ContainerStart")
    commands+=("ContainerStats")
    commands+=("ContainerStop")
    commands+=("ContainerTop")
    commands+=("ContainerUnpause")
    commands+=("ContainerUpdate")
    commands+=("ContainerWait")
    commands+=("PutContainerArchive")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_distribution_DistributionInspect()
{
    last_command="dockerctl_distribution_DistributionInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_distribution()
{
    last_command="dockerctl_distribution"

    command_aliases=()

    commands=()
    commands+=("DistributionInspect")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_exec_ContainerExec()
{
    last_command="dockerctl_exec_ContainerExec"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--containerExecBody.AttachStderr")
    flags+=("--containerExecBody.AttachStdin")
    flags+=("--containerExecBody.AttachStdout")
    flags+=("--containerExecBody.DetachKeys=")
    two_word_flags+=("--containerExecBody.DetachKeys")
    flags+=("--containerExecBody.Privileged")
    flags+=("--containerExecBody.Tty")
    flags+=("--containerExecBody.User=")
    two_word_flags+=("--containerExecBody.User")
    flags+=("--containerExecBody.WorkingDir=")
    two_word_flags+=("--containerExecBody.WorkingDir")
    flags+=("--execConfig=")
    two_word_flags+=("--execConfig")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_exec_ExecInspect()
{
    last_command="dockerctl_exec_ExecInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_exec_ExecResize()
{
    last_command="dockerctl_exec_ExecResize"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--h=")
    two_word_flags+=("--h")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--w=")
    two_word_flags+=("--w")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_exec_ExecStart()
{
    last_command="dockerctl_exec_ExecStart"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--execStartBody.Detach")
    flags+=("--execStartBody.Tty")
    flags+=("--execStartConfig=")
    two_word_flags+=("--execStartConfig")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_exec()
{
    last_command="dockerctl_exec"

    command_aliases=()

    commands=()
    commands+=("ContainerExec")
    commands+=("ExecInspect")
    commands+=("ExecResize")
    commands+=("ExecStart")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_help()
{
    last_command="dockerctl_help"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    has_completion_function=1
    noun_aliases=()
}

_dockerctl_image_BuildPrune()
{
    last_command="dockerctl_image_BuildPrune"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--all")
    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--keep-storage=")
    two_word_flags+=("--keep-storage")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImageBuild()
{
    last_command="dockerctl_image_ImageBuild"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--Content-type=")
    two_word_flags+=("--Content-type")
    flags+=("--X-Registry-Config=")
    two_word_flags+=("--X-Registry-Config")
    flags+=("--buildargs=")
    two_word_flags+=("--buildargs")
    flags+=("--cachefrom=")
    two_word_flags+=("--cachefrom")
    flags+=("--cpuperiod=")
    two_word_flags+=("--cpuperiod")
    flags+=("--cpuquota=")
    two_word_flags+=("--cpuquota")
    flags+=("--cpusetcpus=")
    two_word_flags+=("--cpusetcpus")
    flags+=("--cpushares=")
    two_word_flags+=("--cpushares")
    flags+=("--dockerfile=")
    two_word_flags+=("--dockerfile")
    flags+=("--extrahosts=")
    two_word_flags+=("--extrahosts")
    flags+=("--forcerm")
    flags+=("--labels=")
    two_word_flags+=("--labels")
    flags+=("--memory=")
    two_word_flags+=("--memory")
    flags+=("--memswap=")
    two_word_flags+=("--memswap")
    flags+=("--networkmode=")
    two_word_flags+=("--networkmode")
    flags+=("--nocache")
    flags+=("--outputs=")
    two_word_flags+=("--outputs")
    flags+=("--platform=")
    two_word_flags+=("--platform")
    flags+=("--pull=")
    two_word_flags+=("--pull")
    flags+=("--q")
    flags+=("--remote=")
    two_word_flags+=("--remote")
    flags+=("--rm")
    flags+=("--shmsize=")
    two_word_flags+=("--shmsize")
    flags+=("--squash")
    flags+=("--t=")
    two_word_flags+=("--t")
    flags+=("--target=")
    two_word_flags+=("--target")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImageCommit()
{
    last_command="dockerctl_image_ImageCommit"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--author=")
    two_word_flags+=("--author")
    flags+=("--changes=")
    two_word_flags+=("--changes")
    flags+=("--comment=")
    two_word_flags+=("--comment")
    flags+=("--container=")
    two_word_flags+=("--container")
    flags+=("--containerConfig=")
    two_word_flags+=("--containerConfig")
    flags+=("--containerConfig.ArgsEscaped")
    flags+=("--containerConfig.AttachStderr")
    flags+=("--containerConfig.AttachStdin")
    flags+=("--containerConfig.AttachStdout")
    flags+=("--containerConfig.Domainname=")
    two_word_flags+=("--containerConfig.Domainname")
    flags+=("--containerConfig.Healthcheck.Interval=")
    two_word_flags+=("--containerConfig.Healthcheck.Interval")
    flags+=("--containerConfig.Healthcheck.Retries=")
    two_word_flags+=("--containerConfig.Healthcheck.Retries")
    flags+=("--containerConfig.Healthcheck.StartPeriod=")
    two_word_flags+=("--containerConfig.Healthcheck.StartPeriod")
    flags+=("--containerConfig.Healthcheck.Timeout=")
    two_word_flags+=("--containerConfig.Healthcheck.Timeout")
    flags+=("--containerConfig.Hostname=")
    two_word_flags+=("--containerConfig.Hostname")
    flags+=("--containerConfig.Image=")
    two_word_flags+=("--containerConfig.Image")
    flags+=("--containerConfig.MacAddress=")
    two_word_flags+=("--containerConfig.MacAddress")
    flags+=("--containerConfig.NetworkDisabled")
    flags+=("--containerConfig.OpenStdin")
    flags+=("--containerConfig.StdinOnce")
    flags+=("--containerConfig.StopSignal=")
    two_word_flags+=("--containerConfig.StopSignal")
    flags+=("--containerConfig.StopTimeout=")
    two_word_flags+=("--containerConfig.StopTimeout")
    flags+=("--containerConfig.Tty")
    flags+=("--containerConfig.User=")
    two_word_flags+=("--containerConfig.User")
    flags+=("--containerConfig.WorkingDir=")
    two_word_flags+=("--containerConfig.WorkingDir")
    flags+=("--pause")
    flags+=("--repo=")
    two_word_flags+=("--repo")
    flags+=("--tag=")
    two_word_flags+=("--tag")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImageCreate()
{
    last_command="dockerctl_image_ImageCreate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--X-Registry-Auth=")
    two_word_flags+=("--X-Registry-Auth")
    flags+=("--fromImage=")
    two_word_flags+=("--fromImage")
    flags+=("--fromSrc=")
    two_word_flags+=("--fromSrc")
    flags+=("--inputImage=")
    two_word_flags+=("--inputImage")
    flags+=("--platform=")
    two_word_flags+=("--platform")
    flags+=("--repo=")
    two_word_flags+=("--repo")
    flags+=("--tag=")
    two_word_flags+=("--tag")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImageDelete()
{
    last_command="dockerctl_image_ImageDelete"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--force")
    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--noprune")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImageGet()
{
    last_command="dockerctl_image_ImageGet"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImageGetAll()
{
    last_command="dockerctl_image_ImageGetAll"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImageHistory()
{
    last_command="dockerctl_image_ImageHistory"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImageInspect()
{
    last_command="dockerctl_image_ImageInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImageList()
{
    last_command="dockerctl_image_ImageList"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--all")
    flags+=("--digests")
    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImageLoad()
{
    last_command="dockerctl_image_ImageLoad"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--quiet")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImagePrune()
{
    last_command="dockerctl_image_ImagePrune"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImagePush()
{
    last_command="dockerctl_image_ImagePush"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--X-Registry-Auth=")
    two_word_flags+=("--X-Registry-Auth")
    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--tag=")
    two_word_flags+=("--tag")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImageSearch()
{
    last_command="dockerctl_image_ImageSearch"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--limit=")
    two_word_flags+=("--limit")
    flags+=("--term=")
    two_word_flags+=("--term")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image_ImageTag()
{
    last_command="dockerctl_image_ImageTag"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--repo=")
    two_word_flags+=("--repo")
    flags+=("--tag=")
    two_word_flags+=("--tag")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_image()
{
    last_command="dockerctl_image"

    command_aliases=()

    commands=()
    commands+=("BuildPrune")
    commands+=("ImageBuild")
    commands+=("ImageCommit")
    commands+=("ImageCreate")
    commands+=("ImageDelete")
    commands+=("ImageGet")
    commands+=("ImageGetAll")
    commands+=("ImageHistory")
    commands+=("ImageInspect")
    commands+=("ImageList")
    commands+=("ImageLoad")
    commands+=("ImagePrune")
    commands+=("ImagePush")
    commands+=("ImageSearch")
    commands+=("ImageTag")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_network_NetworkConnect()
{
    last_command="dockerctl_network_NetworkConnect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--container=")
    two_word_flags+=("--container")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--networkConnectBody.Container=")
    two_word_flags+=("--networkConnectBody.Container")
    flags+=("--networkConnectBody.EndpointConfig.EndpointID=")
    two_word_flags+=("--networkConnectBody.EndpointConfig.EndpointID")
    flags+=("--networkConnectBody.EndpointConfig.Gateway=")
    two_word_flags+=("--networkConnectBody.EndpointConfig.Gateway")
    flags+=("--networkConnectBody.EndpointConfig.GlobalIPv6Address=")
    two_word_flags+=("--networkConnectBody.EndpointConfig.GlobalIPv6Address")
    flags+=("--networkConnectBody.EndpointConfig.GlobalIPv6PrefixLen=")
    two_word_flags+=("--networkConnectBody.EndpointConfig.GlobalIPv6PrefixLen")
    flags+=("--networkConnectBody.EndpointConfig.IPAMConfig.IPv4Address=")
    two_word_flags+=("--networkConnectBody.EndpointConfig.IPAMConfig.IPv4Address")
    flags+=("--networkConnectBody.EndpointConfig.IPAMConfig.IPv6Address=")
    two_word_flags+=("--networkConnectBody.EndpointConfig.IPAMConfig.IPv6Address")
    flags+=("--networkConnectBody.EndpointConfig.IPAddress=")
    two_word_flags+=("--networkConnectBody.EndpointConfig.IPAddress")
    flags+=("--networkConnectBody.EndpointConfig.IPPrefixLen=")
    two_word_flags+=("--networkConnectBody.EndpointConfig.IPPrefixLen")
    flags+=("--networkConnectBody.EndpointConfig.IPv6Gateway=")
    two_word_flags+=("--networkConnectBody.EndpointConfig.IPv6Gateway")
    flags+=("--networkConnectBody.EndpointConfig.MacAddress=")
    two_word_flags+=("--networkConnectBody.EndpointConfig.MacAddress")
    flags+=("--networkConnectBody.EndpointConfig.NetworkID=")
    two_word_flags+=("--networkConnectBody.EndpointConfig.NetworkID")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_network_NetworkCreate()
{
    last_command="dockerctl_network_NetworkCreate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--networkConfig=")
    two_word_flags+=("--networkConfig")
    flags+=("--networkCreateBody.Attachable")
    flags+=("--networkCreateBody.CheckDuplicate")
    flags+=("--networkCreateBody.Driver=")
    two_word_flags+=("--networkCreateBody.Driver")
    flags+=("--networkCreateBody.EnableIPv6")
    flags+=("--networkCreateBody.IPAM.Driver=")
    two_word_flags+=("--networkCreateBody.IPAM.Driver")
    flags+=("--networkCreateBody.Ingress")
    flags+=("--networkCreateBody.Internal")
    flags+=("--networkCreateBody.Name=")
    two_word_flags+=("--networkCreateBody.Name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_network_NetworkDelete()
{
    last_command="dockerctl_network_NetworkDelete"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_network_NetworkDisconnect()
{
    last_command="dockerctl_network_NetworkDisconnect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--container=")
    two_word_flags+=("--container")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--networkDisconnectBody.Container=")
    two_word_flags+=("--networkDisconnectBody.Container")
    flags+=("--networkDisconnectBody.Force")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_network_NetworkInspect()
{
    last_command="dockerctl_network_NetworkInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--scope=")
    two_word_flags+=("--scope")
    flags+=("--verbose")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_network_NetworkList()
{
    last_command="dockerctl_network_NetworkList"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_network_NetworkPrune()
{
    last_command="dockerctl_network_NetworkPrune"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_network()
{
    last_command="dockerctl_network"

    command_aliases=()

    commands=()
    commands+=("NetworkConnect")
    commands+=("NetworkCreate")
    commands+=("NetworkDelete")
    commands+=("NetworkDisconnect")
    commands+=("NetworkInspect")
    commands+=("NetworkList")
    commands+=("NetworkPrune")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_node_NodeDelete()
{
    last_command="dockerctl_node_NodeDelete"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--force")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_node_NodeInspect()
{
    last_command="dockerctl_node_NodeInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_node_NodeList()
{
    last_command="dockerctl_node_NodeList"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_node_NodeUpdate()
{
    last_command="dockerctl_node_NodeUpdate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--body=")
    two_word_flags+=("--body")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--nodeSpec.Availability=")
    two_word_flags+=("--nodeSpec.Availability")
    flags+=("--nodeSpec.Name=")
    two_word_flags+=("--nodeSpec.Name")
    flags+=("--nodeSpec.Role=")
    two_word_flags+=("--nodeSpec.Role")
    flags+=("--version=")
    two_word_flags+=("--version")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_node()
{
    last_command="dockerctl_node"

    command_aliases=()

    commands=()
    commands+=("NodeDelete")
    commands+=("NodeInspect")
    commands+=("NodeList")
    commands+=("NodeUpdate")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_plugin_GetPluginPrivileges()
{
    last_command="dockerctl_plugin_GetPluginPrivileges"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--remote=")
    two_word_flags+=("--remote")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_plugin_PluginCreate()
{
    last_command="dockerctl_plugin_PluginCreate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_plugin_PluginDelete()
{
    last_command="dockerctl_plugin_PluginDelete"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--force")
    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_plugin_PluginDisable()
{
    last_command="dockerctl_plugin_PluginDisable"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_plugin_PluginEnable()
{
    last_command="dockerctl_plugin_PluginEnable"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--timeout=")
    two_word_flags+=("--timeout")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_plugin_PluginInspect()
{
    last_command="dockerctl_plugin_PluginInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_plugin_PluginList()
{
    last_command="dockerctl_plugin_PluginList"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_plugin_PluginPull()
{
    last_command="dockerctl_plugin_PluginPull"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--X-Registry-Auth=")
    two_word_flags+=("--X-Registry-Auth")
    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--remote=")
    two_word_flags+=("--remote")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_plugin_PluginPush()
{
    last_command="dockerctl_plugin_PluginPush"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_plugin_PluginSet()
{
    last_command="dockerctl_plugin_PluginSet"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_plugin_PluginUpgrade()
{
    last_command="dockerctl_plugin_PluginUpgrade"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--X-Registry-Auth=")
    two_word_flags+=("--X-Registry-Auth")
    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--remote=")
    two_word_flags+=("--remote")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_plugin()
{
    last_command="dockerctl_plugin"

    command_aliases=()

    commands=()
    commands+=("GetPluginPrivileges")
    commands+=("PluginCreate")
    commands+=("PluginDelete")
    commands+=("PluginDisable")
    commands+=("PluginEnable")
    commands+=("PluginInspect")
    commands+=("PluginList")
    commands+=("PluginPull")
    commands+=("PluginPush")
    commands+=("PluginSet")
    commands+=("PluginUpgrade")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_secret_SecretCreate()
{
    last_command="dockerctl_secret_SecretCreate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--body=")
    two_word_flags+=("--body")
    flags+=("--secretCreateBody.Data=")
    two_word_flags+=("--secretCreateBody.Data")
    flags+=("--secretCreateBody.Driver.Name=")
    two_word_flags+=("--secretCreateBody.Driver.Name")
    flags+=("--secretCreateBody.Name=")
    two_word_flags+=("--secretCreateBody.Name")
    flags+=("--secretCreateBody.Templating.Name=")
    two_word_flags+=("--secretCreateBody.Templating.Name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_secret_SecretDelete()
{
    last_command="dockerctl_secret_SecretDelete"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_secret_SecretInspect()
{
    last_command="dockerctl_secret_SecretInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_secret_SecretList()
{
    last_command="dockerctl_secret_SecretList"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_secret_SecretUpdate()
{
    last_command="dockerctl_secret_SecretUpdate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--body=")
    two_word_flags+=("--body")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--secretSpec.Data=")
    two_word_flags+=("--secretSpec.Data")
    flags+=("--secretSpec.Driver.Name=")
    two_word_flags+=("--secretSpec.Driver.Name")
    flags+=("--secretSpec.Name=")
    two_word_flags+=("--secretSpec.Name")
    flags+=("--secretSpec.Templating.Name=")
    two_word_flags+=("--secretSpec.Templating.Name")
    flags+=("--version=")
    two_word_flags+=("--version")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_secret()
{
    last_command="dockerctl_secret"

    command_aliases=()

    commands=()
    commands+=("SecretCreate")
    commands+=("SecretDelete")
    commands+=("SecretInspect")
    commands+=("SecretList")
    commands+=("SecretUpdate")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_service_ServiceCreate()
{
    last_command="dockerctl_service_ServiceCreate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--X-Registry-Auth=")
    two_word_flags+=("--X-Registry-Auth")
    flags+=("--body=")
    two_word_flags+=("--body")
    flags+=("--serviceCreateBody.EndpointSpec.Mode=")
    two_word_flags+=("--serviceCreateBody.EndpointSpec.Mode")
    flags+=("--serviceCreateBody.Mode.Replicated.Replicas=")
    two_word_flags+=("--serviceCreateBody.Mode.Replicated.Replicas")
    flags+=("--serviceCreateBody.Name=")
    two_word_flags+=("--serviceCreateBody.Name")
    flags+=("--serviceCreateBody.RollbackConfig.Delay=")
    two_word_flags+=("--serviceCreateBody.RollbackConfig.Delay")
    flags+=("--serviceCreateBody.RollbackConfig.FailureAction=")
    two_word_flags+=("--serviceCreateBody.RollbackConfig.FailureAction")
    flags+=("--serviceCreateBody.RollbackConfig.MaxFailureRatio=")
    two_word_flags+=("--serviceCreateBody.RollbackConfig.MaxFailureRatio")
    flags+=("--serviceCreateBody.RollbackConfig.Monitor=")
    two_word_flags+=("--serviceCreateBody.RollbackConfig.Monitor")
    flags+=("--serviceCreateBody.RollbackConfig.Order=")
    two_word_flags+=("--serviceCreateBody.RollbackConfig.Order")
    flags+=("--serviceCreateBody.RollbackConfig.Parallelism=")
    two_word_flags+=("--serviceCreateBody.RollbackConfig.Parallelism")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Dir=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Dir")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.HealthCheck.Interval=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.HealthCheck.Interval")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.HealthCheck.Retries=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.HealthCheck.Retries")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.HealthCheck.StartPeriod=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.HealthCheck.StartPeriod")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.HealthCheck.Timeout=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.HealthCheck.Timeout")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Hostname=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Hostname")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Image=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Image")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Init")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Isolation=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Isolation")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.OpenStdin")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.CredentialSpec.Config=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.CredentialSpec.Config")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.CredentialSpec.File=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.CredentialSpec.File")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.CredentialSpec.Registry=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.CredentialSpec.Registry")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Disable")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Level=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Level")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Role=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Role")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Type=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Type")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.User=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.User")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.ReadOnly")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.StopGracePeriod=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.StopGracePeriod")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.StopSignal=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.StopSignal")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.TTY")
    flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.User=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ContainerSpec.User")
    flags+=("--serviceCreateBody.TaskTemplate.ForceUpdate=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.ForceUpdate")
    flags+=("--serviceCreateBody.TaskTemplate.LogDriver.Name=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.LogDriver.Name")
    flags+=("--serviceCreateBody.TaskTemplate.NetworkAttachmentSpec.ContainerID=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.NetworkAttachmentSpec.ContainerID")
    flags+=("--serviceCreateBody.TaskTemplate.Placement.MaxReplicas=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.Placement.MaxReplicas")
    flags+=("--serviceCreateBody.TaskTemplate.PluginSpec.Disabled")
    flags+=("--serviceCreateBody.TaskTemplate.PluginSpec.Name=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.PluginSpec.Name")
    flags+=("--serviceCreateBody.TaskTemplate.PluginSpec.Remote=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.PluginSpec.Remote")
    flags+=("--serviceCreateBody.TaskTemplate.Resources.Limits.MemoryBytes=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.Resources.Limits.MemoryBytes")
    flags+=("--serviceCreateBody.TaskTemplate.Resources.Limits.NanoCPUs=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.Resources.Limits.NanoCPUs")
    flags+=("--serviceCreateBody.TaskTemplate.Resources.Reservation.MemoryBytes=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.Resources.Reservation.MemoryBytes")
    flags+=("--serviceCreateBody.TaskTemplate.Resources.Reservation.NanoCPUs=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.Resources.Reservation.NanoCPUs")
    flags+=("--serviceCreateBody.TaskTemplate.RestartPolicy.Condition=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.RestartPolicy.Condition")
    flags+=("--serviceCreateBody.TaskTemplate.RestartPolicy.Delay=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.RestartPolicy.Delay")
    flags+=("--serviceCreateBody.TaskTemplate.RestartPolicy.MaxAttempts=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.RestartPolicy.MaxAttempts")
    flags+=("--serviceCreateBody.TaskTemplate.RestartPolicy.Window=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.RestartPolicy.Window")
    flags+=("--serviceCreateBody.TaskTemplate.Runtime=")
    two_word_flags+=("--serviceCreateBody.TaskTemplate.Runtime")
    flags+=("--serviceCreateBody.UpdateConfig.Delay=")
    two_word_flags+=("--serviceCreateBody.UpdateConfig.Delay")
    flags+=("--serviceCreateBody.UpdateConfig.FailureAction=")
    two_word_flags+=("--serviceCreateBody.UpdateConfig.FailureAction")
    flags+=("--serviceCreateBody.UpdateConfig.MaxFailureRatio=")
    two_word_flags+=("--serviceCreateBody.UpdateConfig.MaxFailureRatio")
    flags+=("--serviceCreateBody.UpdateConfig.Monitor=")
    two_word_flags+=("--serviceCreateBody.UpdateConfig.Monitor")
    flags+=("--serviceCreateBody.UpdateConfig.Order=")
    two_word_flags+=("--serviceCreateBody.UpdateConfig.Order")
    flags+=("--serviceCreateBody.UpdateConfig.Parallelism=")
    two_word_flags+=("--serviceCreateBody.UpdateConfig.Parallelism")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_service_ServiceDelete()
{
    last_command="dockerctl_service_ServiceDelete"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_service_ServiceInspect()
{
    last_command="dockerctl_service_ServiceInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--insertDefaults")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_service_ServiceList()
{
    last_command="dockerctl_service_ServiceList"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--status")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_service_ServiceLogs()
{
    last_command="dockerctl_service_ServiceLogs"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--details")
    flags+=("--follow")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--since=")
    two_word_flags+=("--since")
    flags+=("--stderr")
    flags+=("--stdout")
    flags+=("--tail=")
    two_word_flags+=("--tail")
    flags+=("--timestamps")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_service_ServiceUpdate()
{
    last_command="dockerctl_service_ServiceUpdate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--X-Registry-Auth=")
    two_word_flags+=("--X-Registry-Auth")
    flags+=("--body=")
    two_word_flags+=("--body")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--registryAuthFrom=")
    two_word_flags+=("--registryAuthFrom")
    flags+=("--rollback=")
    two_word_flags+=("--rollback")
    flags+=("--serviceUpdateBody.EndpointSpec.Mode=")
    two_word_flags+=("--serviceUpdateBody.EndpointSpec.Mode")
    flags+=("--serviceUpdateBody.Mode.Replicated.Replicas=")
    two_word_flags+=("--serviceUpdateBody.Mode.Replicated.Replicas")
    flags+=("--serviceUpdateBody.Name=")
    two_word_flags+=("--serviceUpdateBody.Name")
    flags+=("--serviceUpdateBody.RollbackConfig.Delay=")
    two_word_flags+=("--serviceUpdateBody.RollbackConfig.Delay")
    flags+=("--serviceUpdateBody.RollbackConfig.FailureAction=")
    two_word_flags+=("--serviceUpdateBody.RollbackConfig.FailureAction")
    flags+=("--serviceUpdateBody.RollbackConfig.MaxFailureRatio=")
    two_word_flags+=("--serviceUpdateBody.RollbackConfig.MaxFailureRatio")
    flags+=("--serviceUpdateBody.RollbackConfig.Monitor=")
    two_word_flags+=("--serviceUpdateBody.RollbackConfig.Monitor")
    flags+=("--serviceUpdateBody.RollbackConfig.Order=")
    two_word_flags+=("--serviceUpdateBody.RollbackConfig.Order")
    flags+=("--serviceUpdateBody.RollbackConfig.Parallelism=")
    two_word_flags+=("--serviceUpdateBody.RollbackConfig.Parallelism")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Dir=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Dir")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.HealthCheck.Interval=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.HealthCheck.Interval")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.HealthCheck.Retries=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.HealthCheck.Retries")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.HealthCheck.StartPeriod=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.HealthCheck.StartPeriod")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.HealthCheck.Timeout=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.HealthCheck.Timeout")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Hostname=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Hostname")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Image=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Image")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Init")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Isolation=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Isolation")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.OpenStdin")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.CredentialSpec.Config=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.CredentialSpec.Config")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.CredentialSpec.File=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.CredentialSpec.File")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.CredentialSpec.Registry=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.CredentialSpec.Registry")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Disable")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Level=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Level")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Role=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Role")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Type=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.Type")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.User=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.Privileges.SELinuxContext.User")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.ReadOnly")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.StopGracePeriod=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.StopGracePeriod")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.StopSignal=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.StopSignal")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.TTY")
    flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.User=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ContainerSpec.User")
    flags+=("--serviceUpdateBody.TaskTemplate.ForceUpdate=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.ForceUpdate")
    flags+=("--serviceUpdateBody.TaskTemplate.LogDriver.Name=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.LogDriver.Name")
    flags+=("--serviceUpdateBody.TaskTemplate.NetworkAttachmentSpec.ContainerID=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.NetworkAttachmentSpec.ContainerID")
    flags+=("--serviceUpdateBody.TaskTemplate.Placement.MaxReplicas=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.Placement.MaxReplicas")
    flags+=("--serviceUpdateBody.TaskTemplate.PluginSpec.Disabled")
    flags+=("--serviceUpdateBody.TaskTemplate.PluginSpec.Name=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.PluginSpec.Name")
    flags+=("--serviceUpdateBody.TaskTemplate.PluginSpec.Remote=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.PluginSpec.Remote")
    flags+=("--serviceUpdateBody.TaskTemplate.Resources.Limits.MemoryBytes=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.Resources.Limits.MemoryBytes")
    flags+=("--serviceUpdateBody.TaskTemplate.Resources.Limits.NanoCPUs=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.Resources.Limits.NanoCPUs")
    flags+=("--serviceUpdateBody.TaskTemplate.Resources.Reservation.MemoryBytes=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.Resources.Reservation.MemoryBytes")
    flags+=("--serviceUpdateBody.TaskTemplate.Resources.Reservation.NanoCPUs=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.Resources.Reservation.NanoCPUs")
    flags+=("--serviceUpdateBody.TaskTemplate.RestartPolicy.Condition=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.RestartPolicy.Condition")
    flags+=("--serviceUpdateBody.TaskTemplate.RestartPolicy.Delay=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.RestartPolicy.Delay")
    flags+=("--serviceUpdateBody.TaskTemplate.RestartPolicy.MaxAttempts=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.RestartPolicy.MaxAttempts")
    flags+=("--serviceUpdateBody.TaskTemplate.RestartPolicy.Window=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.RestartPolicy.Window")
    flags+=("--serviceUpdateBody.TaskTemplate.Runtime=")
    two_word_flags+=("--serviceUpdateBody.TaskTemplate.Runtime")
    flags+=("--serviceUpdateBody.UpdateConfig.Delay=")
    two_word_flags+=("--serviceUpdateBody.UpdateConfig.Delay")
    flags+=("--serviceUpdateBody.UpdateConfig.FailureAction=")
    two_word_flags+=("--serviceUpdateBody.UpdateConfig.FailureAction")
    flags+=("--serviceUpdateBody.UpdateConfig.MaxFailureRatio=")
    two_word_flags+=("--serviceUpdateBody.UpdateConfig.MaxFailureRatio")
    flags+=("--serviceUpdateBody.UpdateConfig.Monitor=")
    two_word_flags+=("--serviceUpdateBody.UpdateConfig.Monitor")
    flags+=("--serviceUpdateBody.UpdateConfig.Order=")
    two_word_flags+=("--serviceUpdateBody.UpdateConfig.Order")
    flags+=("--serviceUpdateBody.UpdateConfig.Parallelism=")
    two_word_flags+=("--serviceUpdateBody.UpdateConfig.Parallelism")
    flags+=("--version=")
    two_word_flags+=("--version")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_service()
{
    last_command="dockerctl_service"

    command_aliases=()

    commands=()
    commands+=("ServiceCreate")
    commands+=("ServiceDelete")
    commands+=("ServiceInspect")
    commands+=("ServiceList")
    commands+=("ServiceLogs")
    commands+=("ServiceUpdate")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_session_Session()
{
    last_command="dockerctl_session_Session"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_session()
{
    last_command="dockerctl_session"

    command_aliases=()

    commands=()
    commands+=("Session")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_swarm_SwarmInit()
{
    last_command="dockerctl_swarm_SwarmInit"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--body=")
    two_word_flags+=("--body")
    flags+=("--swarmInitBody.AdvertiseAddr=")
    two_word_flags+=("--swarmInitBody.AdvertiseAddr")
    flags+=("--swarmInitBody.DataPathAddr=")
    two_word_flags+=("--swarmInitBody.DataPathAddr")
    flags+=("--swarmInitBody.ForceNewCluster")
    flags+=("--swarmInitBody.ListenAddr=")
    two_word_flags+=("--swarmInitBody.ListenAddr")
    flags+=("--swarmInitBody.Spec.CAConfig.NodeCertExpiry=")
    two_word_flags+=("--swarmInitBody.Spec.CAConfig.NodeCertExpiry")
    flags+=("--swarmInitBody.Spec.CAConfig.SigningCACert=")
    two_word_flags+=("--swarmInitBody.Spec.CAConfig.SigningCACert")
    flags+=("--swarmInitBody.Spec.CAConfig.SigningCAKey=")
    two_word_flags+=("--swarmInitBody.Spec.CAConfig.SigningCAKey")
    flags+=("--swarmInitBody.Spec.Dispatcher.HeartbeatPeriod=")
    two_word_flags+=("--swarmInitBody.Spec.Dispatcher.HeartbeatPeriod")
    flags+=("--swarmInitBody.Spec.EncryptionConfig.AutoLockManagers")
    flags+=("--swarmInitBody.Spec.Name=")
    two_word_flags+=("--swarmInitBody.Spec.Name")
    flags+=("--swarmInitBody.Spec.Orchestration.TaskHistoryRetentionLimit=")
    two_word_flags+=("--swarmInitBody.Spec.Orchestration.TaskHistoryRetentionLimit")
    flags+=("--swarmInitBody.Spec.Raft.ElectionTick=")
    two_word_flags+=("--swarmInitBody.Spec.Raft.ElectionTick")
    flags+=("--swarmInitBody.Spec.Raft.HeartbeatTick=")
    two_word_flags+=("--swarmInitBody.Spec.Raft.HeartbeatTick")
    flags+=("--swarmInitBody.Spec.TaskDefaults.LogDriver.Name=")
    two_word_flags+=("--swarmInitBody.Spec.TaskDefaults.LogDriver.Name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_swarm_SwarmInspect()
{
    last_command="dockerctl_swarm_SwarmInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_swarm_SwarmJoin()
{
    last_command="dockerctl_swarm_SwarmJoin"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--body=")
    two_word_flags+=("--body")
    flags+=("--swarmJoinBody.AdvertiseAddr=")
    two_word_flags+=("--swarmJoinBody.AdvertiseAddr")
    flags+=("--swarmJoinBody.DataPathAddr=")
    two_word_flags+=("--swarmJoinBody.DataPathAddr")
    flags+=("--swarmJoinBody.JoinToken=")
    two_word_flags+=("--swarmJoinBody.JoinToken")
    flags+=("--swarmJoinBody.ListenAddr=")
    two_word_flags+=("--swarmJoinBody.ListenAddr")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_swarm_SwarmLeave()
{
    last_command="dockerctl_swarm_SwarmLeave"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--force")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_swarm_SwarmUnlock()
{
    last_command="dockerctl_swarm_SwarmUnlock"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--body=")
    two_word_flags+=("--body")
    flags+=("--swarmUnlockBody.UnlockKey=")
    two_word_flags+=("--swarmUnlockBody.UnlockKey")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_swarm_SwarmUnlockkey()
{
    last_command="dockerctl_swarm_SwarmUnlockkey"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_swarm_SwarmUpdate()
{
    last_command="dockerctl_swarm_SwarmUpdate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--body=")
    two_word_flags+=("--body")
    flags+=("--rotateManagerToken")
    flags+=("--rotateManagerUnlockKey")
    flags+=("--rotateWorkerToken")
    flags+=("--swarmSpec.CAConfig.NodeCertExpiry=")
    two_word_flags+=("--swarmSpec.CAConfig.NodeCertExpiry")
    flags+=("--swarmSpec.CAConfig.SigningCACert=")
    two_word_flags+=("--swarmSpec.CAConfig.SigningCACert")
    flags+=("--swarmSpec.CAConfig.SigningCAKey=")
    two_word_flags+=("--swarmSpec.CAConfig.SigningCAKey")
    flags+=("--swarmSpec.Dispatcher.HeartbeatPeriod=")
    two_word_flags+=("--swarmSpec.Dispatcher.HeartbeatPeriod")
    flags+=("--swarmSpec.EncryptionConfig.AutoLockManagers")
    flags+=("--swarmSpec.Name=")
    two_word_flags+=("--swarmSpec.Name")
    flags+=("--swarmSpec.Orchestration.TaskHistoryRetentionLimit=")
    two_word_flags+=("--swarmSpec.Orchestration.TaskHistoryRetentionLimit")
    flags+=("--swarmSpec.Raft.ElectionTick=")
    two_word_flags+=("--swarmSpec.Raft.ElectionTick")
    flags+=("--swarmSpec.Raft.HeartbeatTick=")
    two_word_flags+=("--swarmSpec.Raft.HeartbeatTick")
    flags+=("--swarmSpec.TaskDefaults.LogDriver.Name=")
    two_word_flags+=("--swarmSpec.TaskDefaults.LogDriver.Name")
    flags+=("--version=")
    two_word_flags+=("--version")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_swarm()
{
    last_command="dockerctl_swarm"

    command_aliases=()

    commands=()
    commands+=("SwarmInit")
    commands+=("SwarmInspect")
    commands+=("SwarmJoin")
    commands+=("SwarmLeave")
    commands+=("SwarmUnlock")
    commands+=("SwarmUnlockkey")
    commands+=("SwarmUpdate")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_system_SystemAuth()
{
    last_command="dockerctl_system_SystemAuth"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--authConfig=")
    two_word_flags+=("--authConfig")
    flags+=("--authConfig.email=")
    two_word_flags+=("--authConfig.email")
    flags+=("--authConfig.password=")
    two_word_flags+=("--authConfig.password")
    flags+=("--authConfig.serveraddress=")
    two_word_flags+=("--authConfig.serveraddress")
    flags+=("--authConfig.username=")
    two_word_flags+=("--authConfig.username")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_system_SystemDataUsage()
{
    last_command="dockerctl_system_SystemDataUsage"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_system_SystemEvents()
{
    last_command="dockerctl_system_SystemEvents"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--since=")
    two_word_flags+=("--since")
    flags+=("--until=")
    two_word_flags+=("--until")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_system_SystemInfo()
{
    last_command="dockerctl_system_SystemInfo"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_system_SystemPing()
{
    last_command="dockerctl_system_SystemPing"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_system_SystemPingHead()
{
    last_command="dockerctl_system_SystemPingHead"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_system_SystemVersion()
{
    last_command="dockerctl_system_SystemVersion"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_system()
{
    last_command="dockerctl_system"

    command_aliases=()

    commands=()
    commands+=("SystemAuth")
    commands+=("SystemDataUsage")
    commands+=("SystemEvents")
    commands+=("SystemInfo")
    commands+=("SystemPing")
    commands+=("SystemPingHead")
    commands+=("SystemVersion")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_task_TaskInspect()
{
    last_command="dockerctl_task_TaskInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_task_TaskList()
{
    last_command="dockerctl_task_TaskList"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_task_TaskLogs()
{
    last_command="dockerctl_task_TaskLogs"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--details")
    flags+=("--follow")
    flags+=("--id=")
    two_word_flags+=("--id")
    flags+=("--since=")
    two_word_flags+=("--since")
    flags+=("--stderr")
    flags+=("--stdout")
    flags+=("--tail=")
    two_word_flags+=("--tail")
    flags+=("--timestamps")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_task()
{
    last_command="dockerctl_task"

    command_aliases=()

    commands=()
    commands+=("TaskInspect")
    commands+=("TaskList")
    commands+=("TaskLogs")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_volume_VolumeCreate()
{
    last_command="dockerctl_volume_VolumeCreate"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--volumeConfig=")
    two_word_flags+=("--volumeConfig")
    flags+=("--volumeCreateBody.Driver=")
    two_word_flags+=("--volumeCreateBody.Driver")
    flags+=("--volumeCreateBody.Name=")
    two_word_flags+=("--volumeCreateBody.Name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_volume_VolumeDelete()
{
    last_command="dockerctl_volume_VolumeDelete"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--force")
    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_volume_VolumeInspect()
{
    last_command="dockerctl_volume_VolumeInspect"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--name=")
    two_word_flags+=("--name")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_volume_VolumeList()
{
    last_command="dockerctl_volume_VolumeList"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_volume_VolumePrune()
{
    last_command="dockerctl_volume_VolumePrune"

    command_aliases=()

    commands=()

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--filters=")
    two_word_flags+=("--filters")
    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_volume()
{
    last_command="dockerctl_volume"

    command_aliases=()

    commands=()
    commands+=("VolumeCreate")
    commands+=("VolumeDelete")
    commands+=("VolumeInspect")
    commands+=("VolumeList")
    commands+=("VolumePrune")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

_dockerctl_root_command()
{
    last_command="dockerctl"

    command_aliases=()

    commands=()
    commands+=("completion")
    commands+=("config")
    commands+=("container")
    commands+=("distribution")
    commands+=("exec")
    commands+=("help")
    commands+=("image")
    commands+=("network")
    commands+=("node")
    commands+=("plugin")
    commands+=("secret")
    commands+=("service")
    commands+=("session")
    commands+=("swarm")
    commands+=("system")
    commands+=("task")
    commands+=("volume")

    flags=()
    two_word_flags=()
    local_nonpersistent_flags=()
    flags_with_completion=()
    flags_completion=()

    flags+=("--config=")
    two_word_flags+=("--config")
    flags+=("--debug")
    flags+=("--hostname=")
    two_word_flags+=("--hostname")
    flags+=("--scheme=")
    two_word_flags+=("--scheme")

    must_have_one_flag=()
    must_have_one_noun=()
    noun_aliases=()
}

__start_dockerctl()
{
    local cur prev words cword
    declare -A flaghash 2>/dev/null || :
    declare -A aliashash 2>/dev/null || :
    if declare -F _init_completion >/dev/null 2>&1; then
        _init_completion -s || return
    else
        __dockerctl_init_completion -n "=" || return
    fi

    local c=0
    local flags=()
    local two_word_flags=()
    local local_nonpersistent_flags=()
    local flags_with_completion=()
    local flags_completion=()
    local commands=("dockerctl")
    local must_have_one_flag=()
    local must_have_one_noun=()
    local has_completion_function
    local last_command
    local nouns=()

    __dockerctl_handle_word
}

if [[ $(type -t compopt) = "builtin" ]]; then
    complete -o default -F __start_dockerctl dockerctl
else
    complete -o default -o nospace -F __start_dockerctl dockerctl
fi

# ex: ts=4 sw=4 et filetype=sh
