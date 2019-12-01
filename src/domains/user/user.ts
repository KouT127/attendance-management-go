export class User extends Document {
    username: string;
    email: string;
    imageUrl: string;
    shouldEdit: boolean;

    constructor(username: string, email: string, imageUrl: string, shouldEdit = true) {
        super();
        this.username = username;
        this.email = email;
        this.imageUrl = imageUrl;
        this.shouldEdit = shouldEdit;
    }

    static initializeUser(): User {
        return new User('', '', '', true)
    }
}
