## Installation
### From go install
```shell
go install github.com/xh-dev-go/split-string
```

### From bat window
```shell
git clone https://github.com/xh-dev-go/split-string.git
cd split-string
build.bat
```

## Usage
```cmd
# default delimiter is ;
echo %PATH% | split

# same as default
echo %PATH% | split -delimiter ;

# take the first char as separator
echo %PATH% | split -delimiter ;:
```
