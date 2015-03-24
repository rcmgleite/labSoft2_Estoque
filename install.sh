clear
# 1) clean up previous installation
echo '> Cleaning up previous installation...'
initialDir=$(pwd)
echo '> Entering $GOPATH/bin dir at: '$GOPATH
cd $GOPATH/bin
rm -rf views/ labEngSoft_Estoque
cd $initialDir
echo '>> Done!'

# 2) creating database
echo '> Do you want to create/recreate database? (y/n)'
read answer1
if [[ $answer1 = "y" ]]
  then
    sqlite3 estoque.db < createDatabase.sql # create database estoque.db
    cd $GOPATH/bin # cd to $GOPATH/bin
    rm -rf estoque.db # remove old db
    cd $initialDir # return to initialDir
    mv estoque.db $GOPATH/bin # move new db to $GOPATH/bin
    echo '>> Done!'
else
  echo '> Skipping Database creation...'
fi

# 3) Execute go install
echo '> Preparing to execute "go install"...'
go install
echo '>> Done!'

echo '>> Installation Complete!'

# 4) Run server
echo '> Do you want to run the server right now? (y/n)'
read answer2
if [[ $answer2 = "y" ]]
  then
  echo '>> Executing server'
  cd $GOPATH/bin/
  ./labEngSoft_Estoque
else
  echo '> Exiting...'
fi
