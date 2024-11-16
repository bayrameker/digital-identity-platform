import Message, { IMessage } from '../models/message';
import socketHandler from '../sockets/socketHandler';

export const sendMessage = async (messageData: IMessage) => {
    const message = new Message(messageData);
    await message.save();

    // Mesajı alıcıya ilet
    socketHandler.emitToUser(message.receiverId, 'message', message);
};
