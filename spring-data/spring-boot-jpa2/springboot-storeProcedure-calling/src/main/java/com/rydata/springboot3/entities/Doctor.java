package com.rydata.springboot3.entities;

import jakarta.persistence.AttributeOverride;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Table;

import com.rydata.springboot3.commons.PersonInfo;

@Entity
@Table(name="doctors")
@AttributeOverride(name = "entrydate", column = @Column(name ="joiningdate"))
public class Doctor extends PersonInfo {
    
    private String roomno;
    private String specialization;

    public String getRoomno() {
        return roomno;
    }
    public void setRoomno(String roomno) {
        this.roomno = roomno;
    }
    public String getSpecialization() {
        return specialization;
    }
    public void setSpecialization(String specialization) {
        this.specialization = specialization;
    }

    
}
