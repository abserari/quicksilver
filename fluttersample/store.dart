import 'dart:indexed_db';

int InsertFeild(String name,int i,num n) {

Database.Transaction(databaseName) async{
  int id = await databaseName.rawInsert(
    'INSERT INTO Test(name, value, num) VALUES(?,?,?)',
    [name,i,n];
  )
}
}
InsertFeild('ydr',4,1999.05)
InsertFeild('x',4,392.05)
InsertFeild('d',4,5.05)
InsertFeild('f',4,1999.05)
