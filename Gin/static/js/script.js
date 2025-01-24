function changeMessage() {
   const messageElement = document.getElementById('message');
   const currentMessage = messageElement.textContent;

   const alternateMessages = [
    "Hello, World!",
    "Goodbye, World!",
    "Hello, desde aquí!"
   ];

   let newMessage;

   do {
       newMessage = alternateMessages[Math.floor(Math.random() * alternateMessages.length)];
   } while (newMessage === currentMessage);

   messageElement.textContent = newMessage;
}