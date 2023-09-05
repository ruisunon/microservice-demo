package com.rydata.springboot3.entities;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Inheritance;
import jakarta.persistence.InheritanceType;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;

import com.fasterxml.jackson.annotation.JsonIgnore;

@Entity
@Inheritance(strategy = InheritanceType.TABLE_PER_CLASS)
public class Vehicle {
    
    @Id
    @GeneratedValue(strategy = GenerationType.SEQUENCE)
    private int id;

    private String modelname;
    private String model;
    private String fueltype;
    private String mileage;
    private String classification;

    private int manufacturer_id;

    @JsonIgnore
    @ManyToOne
    @JoinColumn(name="manufacturer_id",insertable = false, updatable = false)
    private Manufacturer manufacturer;

    public int getId() {
        return id;
    }
    public int getManufacturer_id() {
        return manufacturer_id;
    }
    public void setManufacturer_id(int manufacturer_id) {
        this.manufacturer_id = manufacturer_id;
    }
    public void setId(int id) {
        this.id = id;
    }
    public String getModelname() {
        return modelname;
    }
    public void setModelname(String modelname) {
        this.modelname = modelname;
    }
    public String getModel() {
        return model;
    }
    public void setModel(String model) {
        this.model = model;
    }
    public String getFueltype() {
        return fueltype;
    }
    public void setFueltype(String fueltype) {
        this.fueltype = fueltype;
    }
    public String getMileage() {
        return mileage;
    }
    public void setMileage(String mileage) {
        this.mileage = mileage;
    }
    public String getClassification() {
        return classification;
    }
    public void setClassification(String classification) {
        this.classification = classification;
    }
    public Manufacturer getManufacturer() {
        return manufacturer;
    }
    public void setManufacturer(Manufacturer manufacturer) {
        this.manufacturer = manufacturer;
    }

    

    
}
