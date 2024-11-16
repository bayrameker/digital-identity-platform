export default {
    port: process.env.PORT || 3000,
    jwtSecret: process.env.JWT_SECRET || 'your_jwt_secret',
    mongoURI: process.env.MONGO_URI || 'mongodb://localhost:27017/notification_service',
    redisHost: process.env.REDIS_HOST || 'localhost',
    redisPort: process.env.REDIS_PORT || 6379,
    rabbitMQ: {
        url: process.env.RABBITMQ_URL || 'amqp://localhost',
        queue: process.env.RABBITMQ_QUEUE || 'notification_queue',
    },
    smtp: {
        host: process.env.SMTP_HOST || 'smtp.mailtrap.io',
        port: process.env.SMTP_PORT || 2525,
        user: process.env.SMTP_USER || '',
        pass: process.env.SMTP_PASS || '',
    },
};
