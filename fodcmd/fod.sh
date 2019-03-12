# Created by fodev.org
function fod(){
    case $1 in
        "--enable" | "-e")
            export http_proxy=fodev.org:8118/
            export https_proxy=fodev.org:8118/
            export no_proxy="localhost,127.0.0.1,localaddress,.localdomain.com"
            export HTTP_PROXY=fodev.org:8118/
            export HTTPS_PROXY=fodev.org:8118/
            export NO_PROXY="localhost,127.0.0.1,localaddress,.localdomain.com"
            echo "enable fod proxy !"
        ;;
        "--disable" | "-d")
            unset http_proxy
            unset https_proxy
            unset no_proxy
            unset HTTP_PROXY
            unset HTTPS_PROXY
            unset NO_PROXY
            echo "disable fod proxy !"
        ;;
        *)
            echo "Usage : fod [-e | --enable] [-d | --disable]"
            echo "Example : "
            echo "  fod --enable for enable fod proxy "
            echo "  fod --disable for disable fod proxy "
        ;;
    esac
}