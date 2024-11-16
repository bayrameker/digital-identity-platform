import nodemailer from 'nodemailer';
import config from '../config';

const transporter = nodemailer.createTransport({
    host: config.get('smtp.host'),
    port: config.get('smtp.port'),
    auth: {
        user: config.get('smtp.user'),
        pass: config.get('smtp.pass'),
    },
});

export const sendEmail = async (userId: string, message: string) => {
    // Kullanıcının e-posta adresini Kullanıcı Profil Servisi'nden alın (örnek amaçlı sabit bir adres kullanıyoruz)
    const userEmail = `${userId}@example.com`;

    const mailOptions = {
        from: 'no-reply@example.com',
        to: userEmail,
        subject: 'Yeni Bildiriminiz Var',
        text: message,
    };

    await transporter.sendMail(mailOptions);
};
