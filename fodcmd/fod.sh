# Created by fodev.org
function fod(){
    case $1 in
        "--enable" | "-e")
            export http_proxy=http://fodev.org:8118/
            export https_proxy=http://fodev.org:8118/
            export no_proxy="localhost,127.0.0.1,localaddress,.localdomain.com"
            export HTTP_PROXY=http://fodev.org:8118/
            export HTTPS_PROXY=http://fodev.org:8118/
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
        "--status" | "-s")
            if [[ $http_proxy == 'http://fodev.org:8118/' ]]
            then
                echo 'fod proxy is ENABLED !'
            else
                echo 'fod proxy is DISABLED !'
            fi
        ;;
        *)
            echo "Usage : fod [-e | --enable] [-d | --disable] [-s | --status]"
            echo "Example : "
            echo "  fod --enable to enable fod proxy "
            echo "  fod --disable to disable fod proxy "
            echo "  fod --status to check fod proxy status "
        ;;
    esac
}
