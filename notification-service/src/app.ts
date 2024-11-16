import express from 'express';
import bodyParser from 'body-parser';
import mongoose from 'mongoose';
import config from './config';
import logger from './utils/logger';
import { authMiddleware } from './middlewares/authMiddleware';
import { sendNotification } from './controllers/notificationController';
import { sendMessageController } from './controllers/messageController';

const app = express();

// Middleware'ler
app.use(bodyParser.json());
app.use(authMiddleware);

// Rotalar
app.post('/api/v1/notifications', sendNotification);
app.post('/api/v1/messages', sendMessageController);

// Veritabanına bağlan
mongoose
    .connect(config.get('mongoURI'), { useNewUrlParser: true, useUnifiedTopology: true })
    .then(() => logger.info('MongoDB connected'))
    .catch((err) => logger.error('MongoDB connection error:', err));

export default app;
