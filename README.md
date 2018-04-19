
1. To install mysql in Ubuntu run this:
```
    sudo apt-get update
    sudo apt-get install mysql-server
    mysql_secure_installation

```

2. Check that MySQL is running, execute this command:
```
systemctl status mysql.service
```
If the result shows active (running) then go to next step, otherwise run this command:
```
systemctl start mysql
```

3. Install the tool MySQL workbench , to interact easier with the database:
```
sudo apt install mysql-workbench
```

4. Run MySQL Workbench
```
mysql-workbench
```
Login as root (same password that you entered on step 2)

5. Create a database and use it:
```
CREATE DATABASE recommendations;
USE recommendations;
```

6. Create a schema named "data" with collation of Spanish characters
```
CREATE SCHEMA `data` DEFAULT CHARACTER SET latin1 COLLATE latin1_spanish_ci ;
```

7. Create the table for properties:

```
CREATE TABLE data.realstateproperties (
Id BIGINT,
Type VARCHAR(100),
Business VARCHAR(100),
Sector VARCHAR(100) ,
Value BIGINT ,
City VARCHAR(100) ,
Address VARCHAR(100) ,
Area INT,
Bedrooms INT,
Bathrooms INT,
Parking TINYINT,
StorageRoom TINYINT,
OpenKitchen TINYINT,
Pool TINYINT,
Gym TINYINT,
SocialRoom TINYINT,
CommonAreas TINYINT,
PetFriendly TINYINT,
EcoRoutes TINYINT,
SportFields TINYINT,
WetArea TINYINT,
FrontDeskSecurity TINYINT,
ChildrenPark TINYINT,
ConstructionAge INT,
NearMetro TINYINT,
NearMalls TINYINT,
NearParks TINYINT,
NearMainStreets TINYINT,
NearHospitals TINYINT,
NearRestaurants TINYINT
);

```
