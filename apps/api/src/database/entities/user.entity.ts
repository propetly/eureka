import { Entity, Column, OneToMany } from 'typeorm';
import { BaseEntity } from "~/common/base.entity";
import { Agency } from "./agency.entity";

@Entity()
export class User extends BaseEntity {
    @Column()
    name: string;

    @OneToMany(() => Agency, (agency) => agency.owner)
    agencies: Agency[];
}
