package com.rydata.springboot3.entities.vehicle;

import com.rydata.springboot3.entities.Vehicle;
import jakarta.persistence.Entity;

@Entity

public class Motorbike extends Vehicle {
    
    private String pillionseat;
    private String passengerfootrest;
    private String saddleheight;
    
    public String getPillionseat() {
        return pillionseat;
    }
    public void setPillionseat(String pillionseat) {
        this.pillionseat = pillionseat;
    }
    public String getPassengerfootrest() {
        return passengerfootrest;
    }
    public void setPassengerfootrest(String passengerfootrest) {
        this.passengerfootrest = passengerfootrest;
    }
    public String getSaddleheight() {
        return saddleheight;
    }
    public void setSaddleheight(String saddleheight) {
        this.saddleheight = saddleheight;
    }

    
}
