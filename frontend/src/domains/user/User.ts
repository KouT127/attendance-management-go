export class User extends Document {
    id: string;
    name: string;
    email: string;
    imageUrl: string;
    shouldEdit: boolean;

    constructor(id: string, name: string, email: string, imageUrl: string, shouldEdit = true) {
        super();
        this.id = id;
        this.name = name;
        this.email = email;
        this.imageUrl = imageUrl;
        this.shouldEdit = shouldEdit;
    }

    static initializeUser(userId: string): User {
        return new User(userId, '', '', '', true)
    }

    toJson() {
        return {
            user_name: this.name,
            email: this.email,
            imageUrl: this.imageUrl,
        }
    }
}
