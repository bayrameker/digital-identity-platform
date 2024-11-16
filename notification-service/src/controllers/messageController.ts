import { Request, Response } from 'express';
import { sendMessage } from '../services/messageService';

export const sendMessageController = async (req: Request, res: Response) => {
    try {
        const { senderId, receiverId, content } = req.body;
        await sendMessage({ senderId, receiverId, content });
        res.status(200).json({ message: 'Message sent' });
    } catch (error) {
        res.status(500).json({ error: error.message });
    }
};
