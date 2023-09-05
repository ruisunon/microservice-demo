package com.rydata.springboot3.controllers;

import java.util.List;
import java.util.Set;

import com.rydata.springboot3.Pojos.CityRequest;
import com.rydata.springboot3.Pojos.CourseRequest;
import com.rydata.springboot3.entities.City;
import com.rydata.springboot3.entities.Country;
import com.rydata.springboot3.entities.Course;
import com.rydata.springboot3.services.CityService;
import com.rydata.springboot3.services.CourseService;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HomeController {
    
    @Autowired
    CityService cityService;
    @Autowired
    CourseService courseService;


    @GetMapping("sayhello")
    public String sayHello(){
        return "Hello User";
    }

    @GetMapping("getcities")
    public List<City> getCities(){
        return cityService.getCities();
    }

    @PostMapping("addcity")
    public City SaveCity(@RequestBody City city){
        return cityService.saveCity(city);
    }

    @GetMapping("getcity")
    public City getCity(String cityname){
        return cityService.getCity(cityname);
    }

    @PostMapping("savecity")
    public City addCity(@RequestBody CityRequest cityrequest){
        return cityService.addCity(cityrequest);
    }

    @PostMapping("addcourse")
    public Course addCourse(@RequestBody CourseRequest courseRequest){
        return courseService.addCourseWithContents(courseRequest);
    }

    @GetMapping("countrystartswith")
    public List<Country> countryStartsWith(@RequestParam String countryname){
        return cityService.findByCountryNameStartsWithOrderByPopulation(countryname);
    }
    
    @GetMapping("getcountries")
    public List<Country> getAllCountries(){
        return cityService.getAllCountries();
    }

    @GetMapping("getcountrycontaining")
    public List<Country> getCountryContaining(@RequestParam String substring){
        return cityService.getCountryContaining(substring);
    }

    @GetMapping("getcountry")
    public Country getCountry(@RequestParam int id){
        return cityService.getCountry(id);
    }

    @GetMapping("getcountrybyname")
    public List<Country> getCountryByName(@RequestParam String prefix){
        return cityService.getCountryByName(prefix);
    }

    @GetMapping("getcountrybynameandpop")
    public List<Object[]> getCountryByNameandPop(@RequestParam String prefix, @RequestParam long population){
        return cityService.getCountryByNameandPop(prefix, population);
    }
    
    @GetMapping("getcountrybyids")
    public List<Country> getCountryByIds(@RequestBody Set<Integer> ids){
        return cityService.getCountryByIds(ids);
    }
    
}
