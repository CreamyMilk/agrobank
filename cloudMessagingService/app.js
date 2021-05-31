var serviceAccount = require("./bb.json");
const express = require("express");
const cors = require("cors");
let admin = require('firebase-admin');
require("dotenv").config();

const app = express();
app.use(cors());
app.use(express.json());

let port = process.env.PORT || 8081;

admin.initializeApp({
  credential: admin.credential.cert(serviceAccount),
});


function sendRegistrationData(topicName,role){
    const message = {
        data: {
            type:"role",
            content:role 
        },
         notification:{
            title:"Registration Complete",
            body:role,
        },
      "android":{
       "notification":{
         "icon":"stock_ticker_update",
         "color":"#7e55c3"
       }
     },
        topic: topicName,
    };

    admin
    .messaging()
    .send(message,)
    .then((response) => {
        console.log("Successfully Registration message:", response);
    })
    .catch((error) => {
        console.log("Error sending message:", error);
    });
}
function sendToTopic(topicName,messageTitle,messageDescription,messageType){
    const message = {
        data: {
            type:`${messageType}`,
            content: "A new weather warning has been created!",
        },
         notification:{
            title:messageTitle,
            body:messageDescription,
           image:"https://i.giphy.com/media/14SAx6S02Io1ThOlOY/giphy.webp",
        },
      "android":{
       "notification":{
         "icon":"stock_ticker_update",
         "color":"#7e55c3"
       }
     },
        topic: topicName,
    };

    admin
    .messaging()
    .send(message,)
    .then((response) => {
        console.log("Successfully sent message:", response);
    })
    .catch((error) => {
        console.log("Error sending message:", error);
    });
}
app.get('/',(req,res)=>{    
    res.json({me:"Happy"})
})

app.post("/notifytopic",(req,res)=>{
 let topic = req.body.topic;
 let title = req.body.title;
 let extra = req.body.extra;
 let mtype = req.body.mtype;
sendToTopic(topic,title,extra,mtype);
  res.json({
   status:"0",
   message: "Nofication Sent succesfully",
  })});

app.post("/notifyregistration",(req,res)=>{
 console.table(req.body)
 let topicName = req.body.topic;
 let role = req.body.role;

 sendRegistrationData(topicName,role)
  res.json({
   status:"0",
   message: "Nofication Sent succesfully",
  })});

app.use((req,res,next)=>{
    res.json({status:"We are healthy"})
})
app.listen(port,()=>{
    sendToTopic("pppppppppp","You Have received funds","From James Kamau \n Receipt No:10101010101")
    console.log("FCM",port)
})
