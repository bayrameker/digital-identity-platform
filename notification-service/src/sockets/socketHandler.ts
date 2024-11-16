import { Server } from 'socket.io';

class SocketHandler {
    private io: Server;

    public initialize(server: any) {
        this.io = new Server(server, {
            cors: {
                origin: '*',
            },
        });

        this.io.on('connection', (socket) => {
            console.log('User connected:', socket.id);

            socket.on('join', (userId: string) => {
                socket.join(userId);
            });

            socket.on('disconnect', () => {
                console.log('User disconnected:', socket.id);
            });
        });
    }

    public emitToUser(userId: string, event: string, data: any) {
        this.io.to(userId).emit(event, data);
    }
}

export default new SocketHandler();
