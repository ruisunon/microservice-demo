package com.rydata.springboot3.repositories;

import com.rydata.springboot3.entities.Vehicle;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface VehicleRepository extends JpaRepository<Vehicle,Integer>{
    
}
