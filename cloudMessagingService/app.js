var serviceAccount = require("./cc.json");
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

function sendOrderData(orderObject) {
    const { Topic,
        ProductName,
        Quantity,
        Amount } = orderObject
    const message = {
        data: {
            type: "order",
            prodname: ProductName,
            amount: Amount,
            quantity: Quantity
        },
        notification: {
            title: `Order for ${ProductName} Recieved`,
            body: `Ksh.${Amount}`,
        },
        "android": {
            "notification": {
                "icon": "stock_ticker_update",
                "color": "#7e55c3"
            }
        },
        topic: Topic,
    };

    admin
        .messaging()
        .send(message,)
        .then((response) => {
            console.log("Successfully Deposit message:", response);
        })
        .catch((error) => {
            console.log("Error sending Deposit message:", error);
        });
}




function sendDepositData(topicName, amount) {
    const message = {
        data: {
            type: "deposit",
            amount: amount
        },
        notification: {
            title: "Deposit Recieved",
            body: `Ksh.${amount}`,
        },
        "android": {
            "notification": {
                "icon": "stock_ticker_update",
                "color": "#7e55c3"
            }
        },
        topic: topicName,
    };

    admin
        .messaging()
        .send(message,)
        .then((response) => {
            console.log("Successfully Deposit message:", response);
        })
        .catch((error) => {
            console.log("Error sending Deposit message:", error);
        });
}

function sendRegistrationData(topicName, role) {
    const message = {
        data: {
            type: "role",
            content: role
        },
        notification: {
            title: "Registration Complete",
            body: role,
        },
        "android": {
            "notification": {
                "icon": "stock_ticker_update",
                "color": "#7e55c3"
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
function sendToTopic(topicName, messageTitle, messageDescription, messageType) {
    const message = {
        data: {
            type: `${messageType}`,
            content: "A new weather warning has been created!",
        },
        notification: {
            title: messageTitle,
            body: messageDescription,
        },
        "android": {
            "notification": {
                "icon": "stock_ticker_update",
                "color": "#7e55c3"
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
app.get('/', (req, res) => {
    res.json({ me: "Happy" })
})

app.post("/notifytopic", (req, res) => {
    let topic = req.body.topic;
    let title = req.body.title;
    let extra = req.body.extra;
    let mtype = req.body.mtype;
    sendToTopic(topic, title, extra, mtype);
    res.json({
        status: "0",
        message: "Nofication Sent succesfully",
    })
});

app.post("/notifyregistration", (req, res) => {
    console.table(req.body)
    let topicName = req.body.topic;
    let role = req.body.role;

    sendRegistrationData(topicName, role)
    res.json({
        status: "0",
        message: "Nofication Sent succesfully",
    })
});



app.post("/depositnotif", (req, res) => {
    console.table(req.body)
    let topicName = req.body.topic;
    let amount = req.body.amount;

    sendDepositData(topicName, amount)
    res.json({
        status: "0",
        message: "Deposit Notification Sent succesfully",
    })
});

app.post("/ordernotif", (req, res) => {
    console.table(req.body)
    let order = {
        Topic: req.body.topic,
        ProductName: req.body.prodname,
        Quantity: req.body.quantity,
        Amount: req.body.amount
    }
    sendOrderData(order)
    res.json({
        status: "0",
        message: "Order Placement Worked",
    })
});

app.use((req, res, next) => {
    res.json({ status: "We are healthy" })
})
app.listen(port, () => {
    sendToTopic("all", "You Have received funds", "From James Kamau \n Receipt No:10101010101")
    console.log("FCM", port)
})
