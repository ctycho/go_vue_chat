<template>
    <div class="container">
        <div class="block">
            <div class="title">My Chat</div>
            <div v-if="!messages.length" class="empty">No messages here yet...</div>
            <div id="content" class="content">
                <div class="messages">
                    <div v-for="item of messages">
                        <MessageItem :item="item" :myID="myID" />
                    </div>
                </div>
            </div>
        </div>
        <div class="interact">
            <input type="text" class="input" v-model="text" v-on:keyup.enter="sendMessage()"
                placeholder="Text a message...">
            <button class="button" v-on:click="sendMessage()">Send</button>
        </div>
    </div>
</template>

<script>
import MessageItem from '../components/MessageItem.vue'

export default {
    name: 'HomeView',
    components: {
        MessageItem
    },
    data() {
        return {
            socket: null,
            messages: [],
            text: '',
            myID: ''
        }
    },
    methods: {
        sendMessage() {
            if (!this.text) return
            this.socket.send(this.text)
            this.text = ''
        },
        getTime() {
            const d = new Date()
            let hours = d.getHours()
            let min = d.getMinutes()
            return `${hours < 10 ? `0${hours}` : hours}:${min < 10 ? `0${min}` : min}`
        }
    },
    created() {
        this.socket = new WebSocket("ws://localhost:9000/ws")

        this.socket.onopen = () => {
            console.log('Successfully')
        }

        this.socket.onmessage = (msg) => {
            // console.log('Message: ', msg)

            let data = JSON.parse(msg.data)
            if (!this.myID) this.myID = data.client
            if (data.type == 1 & !!data.body) {
                this.messages.push({
                    client: data.client,
                    text: data.body,
                    time: this.getTime()
                })
            }

        }

        this.socket.onclose = (event) => {
            console.log('Connection closed: ', event)
        }

        this.socket.onerror = (err) => {
            console.log('Connection error: ', err)
        }
    }
}
</script>

<style scoped>
.container {
    background-color: #e3e3e3;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    grid-row-gap: 30px;
    height: 100vh;
}

.block {
    background: #fff;
    padding: 25px;
    border-radius: 10px;
    width: 45%;
    height: 50%;
    overflow: hidden;
}

.title {
    font-size: 26px;
    font-weight: 600;
    text-align: center;
    margin-bottom: 25px;
}

.content {
    height: calc(100% - 35.5px - 25px);
    overflow-y: scroll;
}

.messages {
}

.interact {
    width: calc(40% + 50px);
    display: grid;
    grid-template-columns: 8fr 2fr;
    grid-column-gap: 15px;
}

.input {
    padding: 10px 20px;
    outline: none;
    border: 1px solid white;
    border-radius: 10px;
}

.button {
    background: #83b083;
    border: 1px solid #83b083;
    border-radius: 10px;
    transition: .3s ease;
}

.button:hover {
    cursor: pointer;
    margin-top: -3px;
    margin-bottom: 3px;
}</style>
