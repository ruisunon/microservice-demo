package com.rydata.springboot3.entities.sms;

import com.rydata.springboot3.entities.Sms;
import jakarta.persistence.DiscriminatorValue;
import jakarta.persistence.Entity;

@Entity
@DiscriminatorValue(value = "not null")
public class Other extends Sms {
    
}
