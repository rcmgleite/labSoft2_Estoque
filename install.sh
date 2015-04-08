clear
# 1) clean up previous installation
echo '> Cleaning up previous installation...'
initialDir=$(pwd)
echo '> Entering $GOPATH/bin dir at: '$GOPATH
cd $GOPATH/bin
rm -rf views/ labSoft2_Estoque
cd $initialDir
echo '>> Done!'

# 2) Execute go install
echo '> Preparing to execute "go install"...'
go install
echo '>> Done!'

echo '>> Installation Complete!'

# 3) creating database if one is not already created
if [[ $1 = "-test" ]]
then
	dbName=estoque_test.db
	export TEST=true
else
	export TEST=false
	dbName=estoque.db
fi

cd $GOPATH/bin
if [[ -f $dbName ]]
then
  echo '> Skipping Database creation...'
	cd $initialDir
else
		cd $initialDir
    sqlite3 $dbName < createDatabase.sql # create database estoque.db
    cd $GOPATH/bin # cd to $GOPATH/bin
    rm -rf $dbName # remove old db
    cd $initialDir # return to initialDir
    mv $dbName $GOPATH/bin # move new db to $GOPATH/bin
    echo '>> Done creating db!'
fi

echo 'Executing server...'
cd $GOPATH/bin
./labSoft2_Estoque
