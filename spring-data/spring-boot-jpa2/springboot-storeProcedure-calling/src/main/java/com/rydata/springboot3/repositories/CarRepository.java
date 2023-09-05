package com.rydata.springboot3.repositories;

import com.rydata.springboot3.entities.vehicle.Car;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface CarRepository extends JpaRepository<Car,Integer>{
    
}
