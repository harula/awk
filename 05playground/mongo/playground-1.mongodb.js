/* global use, db */
// MongoDB Playground
// Use Ctrl+Space inside a snippet or a string literal to trigger completions.

const database = 'mydb';
const collection = 'user';

// Create a new database.
use(database);

// Create a new collection if not exist.
db.createCollection(collection);
db.user.drop();

db.user.insertOne(
  {name:'alice',age:24,sex:'m'}
);

db.user.insertMany([
  {name:'bob',age:23,sex:'f'},
  {name:'cat',age:25,sex:'f'},
]);

db.user.find();

db.user.findOne(
  {name:'cat'}
);

//仅更新文档中的字段
db.user.updateOne(
  {name:'cat'},{'$set':{age:13}}
);

//更新所有name=bob的文档
db.user.updateMany(
  {name:'bob'},{'$set':{age:13}}
);

//整个文档替换
db.user.replaceOne(
  {name:'alice'},{name:'alicee',age:10}
);

//删除所有name=cat的文档
db.user.deleteMany(
  {name:'cat'}
);


//创建普通索引：1表示升序，-1表示降序
db.user.createIndex(
  {age:1}
);

//普通索引，可以插入age重复的数据
db.user.insertMany([
  {name:'dog',age:23,sex:'f'},
  {name:'giraffe',age:25,sex:'f'},
]);


//创建复合索引
db.user.createIndex(
  {age:1,name:-1}
);


//查看索引
db.user.getIndexes();

//查看索引的键
db.user.getIndexKeys();


//创建唯一索引
//之后不可以插入name重复的数据
//如果存在重复的name，则创建唯一索引失败
db.user.createIndex(
  {name:1},{unique:true}
);

//此句将插入失败，因为name冲突
/*db.user.insertMany([
  {name:'dog',age:23,sex:'f'},
  {name:'giraffe',age:25,sex:'f'},
]);*/


//索引执行计划，用来分析查询语句是否走索引等
db.user.find(
  {name:'cat'}
).explain();

//ttl索引
db.user.createIndex(
  {age:1},{expireAfterSecs:1}
)



/*db.user.dropIndex(
  {name:1}
);*/

//db.user.dropIndex("age_1_name_-1");
//db.user.drop();
//db.dropDatabase();


