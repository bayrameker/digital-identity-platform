import express from 'express';
import bodyParser from 'body-parser';
import mongoose from 'mongoose';
import { ApolloServer } from 'apollo-server-express';
import config from './utils/config';
import logger from './utils/logger';
import authRoutes from './routes/authRoutes';
import apiRoutes from './routes/apiRoutes';
import webhookRoutes from './routes/webhookRoutes';
import { errorMiddleware } from './middlewares/errorMiddleware';
import typeDefs from './graphql/schema';
import resolvers from './graphql/resolvers';

const app = express();

// Middleware'ler
app.use(bodyParser.json());

// Rotalar
app.use('/auth', authRoutes);
app.use('/api', apiRoutes);
app.use('/webhook', webhookRoutes);

// Hata yönetimi
app.use(errorMiddleware);

// GraphQL sunucusu
const apolloServer = new ApolloServer({ typeDefs, resolvers });
apolloServer.applyMiddleware({ app, path: '/graphql' });

// Veritabanına bağlan
mongoose
    .connect(config.get('mongoURI'), { useNewUrlParser: true, useUnifiedTopology: true })
    .then(() => logger.info('MongoDB connected'))
    .catch((err) => logger.error('MongoDB connection error:', err));

export default app;
