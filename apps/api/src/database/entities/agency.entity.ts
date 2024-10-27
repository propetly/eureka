import { Entity, Column, JoinColumn, ManyToOne } from 'typeorm';
import { BaseEntity } from "~/common/base.entity";
import { User } from "./user.entity";

@Entity()
export class Agency extends BaseEntity {
    @Column()
    name: string;

    @Column({ unique: true })
    slug: string;

    @ManyToOne(() => User, (user) => user)
    @JoinColumn({ name: "ownerId" })
    owner: User
}
