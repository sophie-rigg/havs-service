db = db.getSiblingDB('havs');

db.createUser({
    user: 'admin',
    pwd: 'password',
    roles: [
        {
            role: 'userAdminAnyDatabase',
            db: 'admin',
        },
    ]
});

db.createCollection('users');
db.createCollection('exposures');
db.createCollection('equipment');

db.users.insert({
    _id: 'ec6b9a1e-6036-46f1-91b4-f67806017fe1',
    name: 'John Doe',
})

db.equipment.insert({
    _id: '55499473-7265-4c2e-806e-56c45dea7477',
    name: 'AirCat - Drill - 4337',
    vibrational_magnitude: 2.1,
})

db.equipment.insert({
    _id: '12fe6eb5-04a4-459b-a4b0-15cda0baefb4',
    name: 'JCB - Hydraulic Breaker - CEJCBHM25',
    vibrational_magnitude: 4.0,
})


