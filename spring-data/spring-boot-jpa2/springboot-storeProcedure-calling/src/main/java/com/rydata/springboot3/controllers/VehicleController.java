package com.rydata.springboot3.controllers;

import java.util.List;

import com.rydata.springboot3.entities.Manufacturer;
import com.rydata.springboot3.entities.Vehicle;
import com.rydata.springboot3.entities.vehicle.Car;
import com.rydata.springboot3.entities.vehicle.Motorbike;
import com.rydata.springboot3.repositories.CarRepository;
import com.rydata.springboot3.repositories.ManufacturerRepository;
import com.rydata.springboot3.repositories.MotorbikeRepository;
import com.rydata.springboot3.repositories.VehicleRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("vehicle/")
public class VehicleController {

    @Autowired
    ManufacturerRepository manufacturerRepository;

    @Autowired
    VehicleRepository vehicleRepository;

    @Autowired
    CarRepository carRepository;

    @Autowired
    MotorbikeRepository motorbikeRepository;

    @PostMapping("addmanufacturers")
    public List<Manufacturer> addManufacturers(@RequestBody List<Manufacturer> manufacturers){

        return manufacturerRepository.saveAll(manufacturers);
    }

    @PostMapping("addvehicles")
    public List<Vehicle> addVehicles(@RequestBody List<Vehicle> vehicles){

        return vehicleRepository.saveAll(vehicles);
    }

    @PostMapping("addcars")
    public List<Car> addCars(@RequestBody List<Car> cars){

        return carRepository.saveAll(cars);
    }

    @PostMapping("addmotorbikes")
    public List<Motorbike> addMotorbikes(@RequestBody List<Motorbike> motorbikes){

        return motorbikeRepository.saveAll(motorbikes);
    }
    

    @GetMapping("getvehicles")
    public List<Vehicle> getVehicle(){
        return vehicleRepository.findAll();
    }

    @GetMapping("getcards")
    public List<Car> getCars(){
        return carRepository.findAll();
    }

    @GetMapping("getmotorbikes")
    public List<Motorbike> getMotorbikes(){
        return motorbikeRepository.findAll();
    }

    @GetMapping("getmanufacturers")
    public List<Manufacturer> getManufacturers(){
        return manufacturerRepository.findAll();
    }
}
