package com.rydata.springboot3.entities.vehicle;

import com.rydata.springboot3.entities.Vehicle;
import jakarta.persistence.Entity;

@Entity

public class Car extends Vehicle {

    private String seatingCapacity;
    private String sunroof;
    public String getSeatingCapacity() {
        return seatingCapacity;
    }
    public void setSeatingCapacity(String seatingCapacity) {
        this.seatingCapacity = seatingCapacity;
    }
    public String getSunroof() {
        return sunroof;
    }
    public void setSunroof(String sunroof) {
        this.sunroof = sunroof;
    }

    
    
}
