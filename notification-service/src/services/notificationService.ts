import Notification, { INotification } from '../models/notification';
import UserPreferences from '../models/userPreferences';
import { sendEmail } from './emailService';
import socketHandler from '../sockets/socketHandler';

export const createNotification = async (notificationData: INotification) => {
    const notification = new Notification(notificationData);
    await notification.save();

    // Kullanıcının bildirim tercihlerine göre işlem yap
    const preferences = await UserPreferences.findOne({ userId: notification.userId });
    if (preferences) {
        if (preferences.pushNotifications) {
            // Gerçek zamanlı bildirim gönder
            socketHandler.emitToUser(notification.userId, 'notification', notification);
        }
        if (preferences.emailNotifications) {
            // E-posta gönder
            sendEmail(notification.userId, notification.message);
        }
    } else {
        // Varsayılan olarak push bildirim gönder
        socketHandler.emitToUser(notification.userId, 'notification', notification);
    }
};
