package com.rydata.springboot3.entities.sms;

import com.rydata.springboot3.entities.Sms;
import jakarta.persistence.Entity;

@Entity

public class Inbox extends Sms {

    private String smstype;

    public String getSmstype() {
        return smstype;
    }

    public void setSmstype(String smstype) {
        this.smstype = smstype;
    }

    

}
