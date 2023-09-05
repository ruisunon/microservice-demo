package com.rydata.springboot3.entities.sms;

import jakarta.persistence.Entity;

import com.rydata.springboot3.entities.Sms;

@Entity

public class Failed  extends Sms {

    private String errormessage;

    public String getErrormessage() {
        return errormessage;
    }

    public void setErrormessage(String errormessage) {
        this.errormessage = errormessage;
    }

    
}
