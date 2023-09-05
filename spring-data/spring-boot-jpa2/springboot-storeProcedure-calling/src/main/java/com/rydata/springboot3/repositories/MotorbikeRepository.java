package com.rydata.springboot3.repositories;

import com.rydata.springboot3.entities.vehicle.Motorbike;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface MotorbikeRepository extends JpaRepository<Motorbike,Integer> {
    
}
