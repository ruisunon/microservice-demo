package com.rydata.springboot3.entities.sms;

import java.time.ZonedDateTime;

import jakarta.persistence.Entity;

import com.rydata.springboot3.entities.Sms;

@Entity
public class Sent extends Sms {
    
    private ZonedDateTime deliveredon;

    public ZonedDateTime getDeliveredon() {
        return deliveredon;
    }

    public void setDeliveredon(ZonedDateTime deliveredon) {
        this.deliveredon = deliveredon;
    }

    
}
