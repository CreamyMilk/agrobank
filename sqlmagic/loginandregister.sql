-- Prototype first

User fills in f<rms he can continue where he/she stopped 

User should be able to submit two photos -ID + Him/her self
- image will be handeld as the user takes photo we upload to server and 
return the url back to the user

User can seek to finalize their registration 
->User sends form progress if all is well 
send back a 201 response indicating

that registration fee

[details,number,images,registration point]
points can be eiter 
- form filling
- Payment 
- Done

User is send to payment page
When he/she initiates a payment we make a temp record
user_id,stkPushCode,amount,notificationID
on callback we fix the table issues like deleting the record or storing the record in the db\
Notify the user of the current progress


This will prompt the app to be opened where we will send a request to the server of the current progress
if he is in done stage then prompt the user to set his username
and password


on succesfull details submittion we send back the 
login payload
- User ID
- User Token
- User Wallet
    -Wallet ID
    -Wallet Balance
- User name
- Most recent notifications (Limit to 10)
- Users Orders (No Limit always show all)
- Users transactions (Limit to 10) -> Load More where necessary (infinte list)

and initate heart beats or establish a web socket connection for realtime updates
MESSAGE_PAYLOAD 
//setUp a handler function for many events
- WalletBalance Change
- New Orders
- Delivery is near
- Hear Beats like every 5 seconds will be a marco polo game for sure
- New Transaction Event

User can interact with the shop
Gets a list of products with hopefully server based filter
and serach the details users have can not always pass around hudge json blobs
LIMIT 20
SEARCH BAR WILL SEND request on each key press
on clear fetch stardard 20 which are some popular options by the users



Interaction 
initate Payemnt towards a certain product
Load Wallet
Withdraw form wallet

CREATE TABLE user_registration (
    userid ,
    idnumber,
    idpic,
    phonenumber,
    userpic,
    email,
    Bussiness/Farm address, -> Cords
    Residence,
    Role,
    stage
)


INSERT INTO user_registration (idnumber,idpic,phonenumber,userpic,email,baddress,residence,role)
VALUES ("1091091","https://static.agro","254797333333",
"https://stattic.agro","exaple@com","-2010,2000","BondoSouth",
"Farmer",0)
// on Sends user to the paymentspage
{
    "status":"0",
    "message":"Your Reg was sucessful"
    "Names":"John Doe",
    "zeroformat":"0797678252"
    "temp_userid":"20202"    
}
 or

// Redirects to login page
{
    "status":"90",
    "message":"Your are already registerd k"
    "Names":"John Doe",
    "phoneNumber":"254797678252"
    "zeroformat":"0797678252"
    "temp_userid":"20202"    
} or

// Show just an error and asks user to kindly review or startover
{
    "status":"-1",
    "message":"Falied to register you due ti 1,2,3",
}


if the above operation is successful 
send user to payments page using the response
CREATE TABLE notifications_table{
    userID,
    walletID,
    fcmToken,
    socketID,    
}



on callback use the following to determine if to move user to next stage
or just delete the record from the table
CREATE TABLE user_payment_limbo(
    userid,
    requestSentTo,
    trackingNo
)

when payment is sucessful redirect  user to login page






