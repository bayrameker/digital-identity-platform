export default {
    port: process.env.PORT || 4000,
    jwtSecret: process.env.JWT_SECRET || 'your_jwt_secret',
    mongoURI: process.env.MONGO_URI || 'mongodb://localhost:27017/integration_service',
    redisHost: process.env.REDIS_HOST || 'localhost',
    redisPort: process.env.REDIS_PORT || 6379,
    oauth: {
        clientID: process.env.OAUTH_CLIENT_ID || 'your_client_id',
        clientSecret: process.env.OAUTH_CLIENT_SECRET || 'your_client_secret',
        callbackURL: process.env.OAUTH_CALLBACK_URL || 'http://localhost:4000/auth/callback',
    },
};
