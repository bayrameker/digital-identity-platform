import http from 'http';
import app from './app';
import config from './config';
import socketHandler from './sockets/socketHandler';
import logger from './utils/logger';

// Sunucuyu oluştur
const server = http.createServer(app);

// Socket.IO'yu başlat
socketHandler.initialize(server);

// Sunucuyu başlat
const port = config.get('port');
server.listen(port, () => {
    logger.info(`Server is running on port ${port}`);
});
