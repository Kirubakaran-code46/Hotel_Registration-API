CREATE TABLE `middleware_req_log` (
  `Id` bigint NOT NULL AUTO_INCREMENT,
  `request_id` varchar(100) NOT NULL,
  `token` varchar(100) NOT NULL,
  `requesteddate` datetime NOT NULL,
  `realip` varchar(240) DEFAULT NULL,
  `forwardedip` varchar(200) DEFAULT NULL,
  `method` varchar(20) DEFAULT NULL,
  `path` longtext,
  `host` varchar(240) DEFAULT NULL,
  `remoteaddr` varchar(240) DEFAULT NULL,
  `header` longtext,
  `body` longtext,
  `endpoint` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`Id`)
)

CREATE TABLE `middleware_resp_log` (
  `Id` bigint NOT NULL AUTO_INCREMENT,
  `request_id` varchar(100) DEFAULT NULL,
  `response` longtext,
  `responseStatus` int DEFAULT NULL,
  `requesteddate` datetime NOT NULL,
  `realip` varchar(240) DEFAULT NULL,
  `forwardedip` varchar(200) DEFAULT NULL,
  `method` varchar(20) DEFAULT NULL,
  `path` longtext,
  `host` varchar(240) DEFAULT NULL,
  `remoteaddr` varchar(240) DEFAULT NULL,
  `header` longtext,
  `body` longtext,
  `endpoint` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`Id`)
)


CREATE TABLE `role_management` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `role` varchar(30) NOT NULL,
  `Description` text,
  `CreatedBy` varchar(20) DEFAULT NULL,
  `createdDate` date DEFAULT NULL,
  `UpdatedBy` varchar(20) DEFAULT NULL,
  `UpdatedDate` date DEFAULT NULL,
  PRIMARY KEY (`Id`)
)

CREATE TABLE `accepted_identity_proofs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `proof` varchar(200) DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
)

CREATE TABLE `basic_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` varchar(100) DEFAULT NULL,
  `Hotel_name` varchar(100) DEFAULT NULL,
  `Property_Type` varchar(50) DEFAULT NULL,
  `Email` varchar(50) DEFAULT NULL,
  `Year_Of_Construction` varchar(4) DEFAULT NULL,
  `mobile_code` varchar(5) DEFAULT NULL,
  `Primary_Mobile` varchar(10) DEFAULT NULL,
  `Secondary_Mobile` varchar(10) DEFAULT NULL,
  `Star_Category` char(1) DEFAULT NULL,
  `Channel_Manageer` varchar(5) DEFAULT NULL,
  `CreatedBy` varchar(20) DEFAULT NULL,
  `createdDate` datetime DEFAULT NULL,
  `UpdatedBy` varchar(20) DEFAULT NULL,
  `UpdatedDate` datetime DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  `Description` text DEFAULT NULL,
  PRIMARY KEY (`id`)
) 

CREATE TABLE `country_code` (
  `Country` varchar(50) NOT NULL,
  `mobile_code` varchar(20) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
)

CREATE TABLE `document_upload` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` varchar(100) DEFAULT NULL,
  `Bank_Name` varchar(100) DEFAULT NULL,
  `Account_Number` varchar(50) DEFAULT NULL,
  `Acc_HolderName` varchar(100) DEFAULT NULL,
  `IFSC_code` varchar(50) DEFAULT NULL,
  `Branch` varchar(50) DEFAULT NULL,
  `GST_Number` varchar(50) DEFAULT NULL,
  `GST_docId` varchar(50) DEFAULT NULL,
  `cancelledCheque_docId` varchar(50) DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  `CreatedBy` varchar(20) DEFAULT NULL,
  `createdDate` datetime DEFAULT NULL,
  `UpdatedBy` varchar(20) DEFAULT NULL,
  `UpdatedDate` datetime DEFAULT NULL,
  `propertyOwnership` varchar(80) DEFAULT NULL,
  `Start_Date` varchar(20) DEFAULT NULL,
  `End_Date` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
)

CREATE TABLE `hotel_cancellation_policy` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `policy` varchar(200) DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) 

CREATE TABLE `hotel_property_types` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `property` varchar(50) DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
)

CREATE TABLE `indian_states` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `State_Code` varchar(50) DEFAULT NULL,
  `State_Name` varchar(100) DEFAULT NULL,
  `Type` varchar(50) DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
)

CREATE TABLE `location_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` varchar(100) DEFAULT NULL,
  `Addr_line1` varchar(200) DEFAULT NULL,
  `Addr_line2` varchar(200) DEFAULT NULL,
  `State` varchar(100) DEFAULT NULL,
  `City` varchar(50) DEFAULT NULL,
  `Zip_Code` varchar(50) DEFAULT NULL,
  `CreatedBy` varchar(20) DEFAULT NULL,
  `createdDate` datetime DEFAULT NULL,
  `UpdatedBy` varchar(20) DEFAULT NULL,
  `updatedDate` datetime DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  `availability_Start_Date` date DEFAULT NULL,
  `availability_End_Date` date DEFAULT NULL,
  PRIMARY KEY (`id`)
)

CREATE TABLE `meals_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` varchar(100) DEFAULT NULL,
  `isOptionalRestaurant` varchar(20) DEFAULT NULL,
  `Meal_Package` varchar(50) DEFAULT NULL,
  `Types_Of_Meals` text DEFAULT NULL,
  `Meal_Rack_Price` varchar(50) DEFAULT NULL,
  `CreatedBy` varchar(20) DEFAULT NULL,
  `createdDate` datetime DEFAULT NULL,
  `UpdatedBy` varchar(20) DEFAULT NULL,
  `UpdatedDate` datetime DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
)

CREATE TABLE `property_details` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` varchar(100) DEFAULT NULL,
  `Facade_docId` varchar(50) DEFAULT NULL,
  `Parking_docId` varchar(50) DEFAULT NULL,
  `Lobby_docId` varchar(50) DEFAULT NULL,
  `Reception_docId` varchar(50) DEFAULT NULL,
  `Corridors_docId` varchar(50) DEFAULT NULL,
  `Lift_docId` varchar(50) DEFAULT NULL,
  `Bathroom_docId` varchar(50) DEFAULT NULL,
  `OtherArea_docId` varchar(50) DEFAULT NULL,
  `PropertyImg_docId` varchar(50) DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  `CreatedBy` varchar(20) DEFAULT NULL,
  `createdDate` datetime DEFAULT NULL,
  `UpdatedBy` varchar(20) DEFAULT NULL,
  `UpdatedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)

CREATE TABLE `property_policiesinfo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` varchar(100) DEFAULT NULL,
  `check_in` varchar(20) DEFAULT NULL,
  `check_out` varchar(20) DEFAULT NULL,
  `checkinout_policy` varchar(10) DEFAULT NULL,
  `CancellationPolicy` varchar(100) DEFAULT NULL,
  `allow_unmarriedCouples` varchar(10) DEFAULT NULL,
  `allow_minor_guest` varchar(10) DEFAULT NULL,
  `allow_onlymale_guests` varchar(10) DEFAULT NULL,
  `allow_smoking` varchar(10) DEFAULT NULL,
  `allow_parties` varchar(10) DEFAULT NULL,
  `allow_invite_guests` varchar(10) DEFAULT NULL,
  `wheelchar_accessible` varchar(10) DEFAULT NULL,
  `allow_pets` varchar(10) DEFAULT NULL,
  `accepted_proofs` text DEFAULT NULL,
  `additional_propertyrules` text DEFAULT NULL,
  `CreatedBy` varchar(20) DEFAULT NULL,
  `createdDate` datetime DEFAULT NULL,
  `UpdatedBy` varchar(20) DEFAULT NULL,
  `UpdatedDate` datetime DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) 

CREATE TABLE `room_amenities` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Types` varchar(50) DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
)

CREATE TABLE `room_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Types` varchar(50) DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
)

CREATE TABLE `room_types` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `RoomType` varchar(100) DEFAULT NULL,
  `NoOfRooms` varchar(20) DEFAULT NULL,
  `RoomView` varchar(100) DEFAULT NULL,
  `RoomSizeUnit` varchar(50) DEFAULT NULL,
  `RoomSize` varchar(30) DEFAULT NULL,
  `MaximumOccupancy` varchar(20) DEFAULT NULL,
  `ExtraBed` varchar(20) DEFAULT NULL,
  `ExtraPersons` varchar(20) DEFAULT NULL,
  `SingleGuestPrice` varchar(20) DEFAULT NULL,
  `DoubleGuestPrice` varchar(20) DEFAULT NULL,
  `TripleGuestPrice` varchar(20) DEFAULT NULL,
  `ExtraAdultCharge` varchar(20) DEFAULT NULL,
  `ChildCharge` varchar(20) DEFAULT NULL,
  `BelowChildCharge` varchar(20) DEFAULT NULL,
  `RoomAmenities` text DEFAULT NULL,
  `SmokingPolicy` varchar(20) DEFAULT NULL,
  `CreatedBy` varchar(20) DEFAULT NULL,
  `createdDate` datetime DEFAULT NULL,
  `UpdatedBy` varchar(20) DEFAULT NULL,
  `UpdatedDate` datetime DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  `Uid` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
)


CREATE TABLE `room_view` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Types` varchar(50) DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) 

CREATE TABLE `utility_types` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` varchar(100) DEFAULT NULL,
  `Bill_Type` varchar(100) DEFAULT NULL,
  `Bill_docId` varchar(50) DEFAULT NULL,
  `isActive` char(1) DEFAULT NULL,
  `CreatedBy` varchar(20) DEFAULT NULL,
  `createdDate` datetime DEFAULT NULL,
  `UpdatedBy` varchar(20) DEFAULT NULL,
  `UpdatedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)