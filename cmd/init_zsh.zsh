#compdef _as as

\builtin unalias as &>/dev/null || \builtin true

function as {
    eval "$(asdfswitcher shell $@)"
}

function _as_versions {
    asdf list ${1}
}

function _as_plugins {
    asdf plugin list
}

function _as {
    local context state state_descr line

    _arguments -C \
        "1:plugin:->plugins" \
        "2:version:->versions"

    case "$state" in
        plugins)
            _values 'plugins' $(_as_plugins)
            ;;
        versions)
            _values 'versions' $(_as_versions ${line})
            ;;
    esac
}

compdef _as as
