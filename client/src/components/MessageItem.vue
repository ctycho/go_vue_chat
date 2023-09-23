<template>
    <div class="message" :class="{ my_message: myID == item.client }">
        <div class="client"><span>{{ item.client }}</span></div>
        <div class="text">{{ item.text }}</div>
        <div class="time">{{ item.time }}</div>
    </div>
</template>

<script>

export default {
    name: 'MessageItem',
    props: ['item', 'myID']
}
</script>

<style scoped>
.message {
    padding: 5px;
    border-radius: 5px;
    background-color: #e3e3e3;
    max-width: 60%;
    width: max-content;
    display: flex;
    grid-column-gap: 7px;
    margin-bottom: 10px;
}

.my_message {
    margin-left: auto;
    background-color: #c0e28d;
}

.message:not(:last-child) {
    /* Doesn't work */
    margin-bottom: 10px;
}

.client {
    display: flex;
    align-items: center;
}

.client span {
    background: #fff;
    border-radius: 50%;
    width: 20px;
    height: 20px;
    font-size: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.text {
    text-align: left;
    font-size: 15px;
}

.time {
    display: flex;
    align-items: end;
    font-size: 10px;
}
</style>