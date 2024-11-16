import { Request, Response } from 'express';
import { createNotification } from '../services/notificationService';

export const sendNotification = async (req: Request, res: Response) => {
    try {
        const { userId, message, type } = req.body;
        await createNotification({ userId, message, type });
        res.status(200).json({ message: 'Notification sent' });
    } catch (error) {
        res.status(500).json({ error: error.message });
    }
};
