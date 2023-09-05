/*
 * MIT License
 *
 * Copyright (c) 2020 Alexander Ungar
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package com.netlight.learning.java.spring.petclinic.monolith.customer

import com.fasterxml.jackson.annotation.JsonFormat
import org.springframework.validation.annotation.Validated
import org.springframework.web.bind.annotation.*
import java.util.*
import javax.persistence.Temporal
import javax.persistence.TemporalType
import javax.validation.constraints.NotBlank

@RestController
class PetResource(val petService: PetService, val ownerService: OwnerService, val petMapper: PetMapper) {
    @PostMapping("pets/types")
    fun addPetType(@RequestParam name: String): PetType {
        return with(PetType(0, name)) {
            petService.addPetType(this)
        }
    }

    @GetMapping("pets/types")
    fun getPetTypes(): List<PetType> = petService.getAllTypes()

    @PostMapping("owners/{ownerId}/pets")
    fun addPet(@PathVariable ownerId: Long, @RequestBody @Validated newPetDTO: NewPetDTO): Pet {
        val petType = petService.getPetTypeById(newPetDTO.petTypeId)
        val owner = ownerService.getOne(ownerId)
        val pet = petMapper.from(newPetDTO, owner, petType)

        return petService.addPet(pet)
    }

    @PutMapping("pets/{petId}")
    fun updatePet(@PathVariable petId: Long, @RequestBody @Validated updateDTO: PetUpdateDTO) {
        with(petService.getPetById(petId)) {
            petMapper.updatePet(this, updateDTO)
            petService.updatePet(this, updateDTO.petTypeId)
        }
    }

    @GetMapping("pets/{petId}")
    fun findPet(@PathVariable petId: Long): Pet = petService.getPetById(petId)

    @GetMapping("pets")
    fun getAllPets(): List<Pet> = petService.getAllPets()
}

data class PetUpdateDTO(
        @NotBlank
        var name: String,
        @Temporal(TemporalType.DATE)
        @JsonFormat(pattern = "yyyy-MM-dd")
        var birthDate: Date,
        var petTypeId: Long
)

data class NewPetDTO(
        @NotBlank
        var name: String,
        @Temporal(TemporalType.DATE)
        @JsonFormat(pattern = "yyyy-MM-dd")
        var birthDate: Date,
        var petTypeId: Long
)