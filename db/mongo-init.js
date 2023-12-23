db = db.getSiblingDB('webapp');
db.createCollection('pokemon');
db.pokemons.insertMany( [
	{ _id: 10, item: "large box", qty: 20 },
	{ _id: 11, item: "small box", qty: 55 },
	{ _id: 12, item: "medium box", qty: 30 }
] );
/*
db.createUser(
  {
    user: "go-app",
    pwd: "password",
    roles: [
      {
        role: "readWrite",
        db: "local"
      }
    ]
  }
);
*/
