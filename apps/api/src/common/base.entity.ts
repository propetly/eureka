import { PrimaryGeneratedColumn, CreateDateColumn, UpdateDateColumn } from 'typeorm';


export abstract class BaseEntity {
    @PrimaryGeneratedColumn()
    id: number;

    @PrimaryGeneratedColumn('uuid')
    uuid: string;

    @CreateDateColumn()
    createdAt: Date;

    @UpdateDateColumn()
    updatedAt: Date;
}
