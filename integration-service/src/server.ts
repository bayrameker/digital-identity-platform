import http from 'http';
import app from './app';
import config from './utils/config';
import logger from './utils/logger';

// Sunucuyu oluştur
const server = http.createServer(app);

// Sunucuyu başlat
const port = config.get('port');
server.listen(port, () => {
    logger.info(`Server is running on port ${port}`);
});
