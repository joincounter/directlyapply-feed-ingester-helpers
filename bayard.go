package helpers

// BayardFilter this will filter out bayard companies
func BayardFilter(job *StandardJob) *StandardJob {
	if StringInSlice(job.Company, bayard) {
		return nil
	}
	return job
}

var bayard = []string{"Teleperformance",
	"Research Medical Center",
	"Overland Park Regional Medical Center",
	"Garden Park Medical Center",
	"Belton Regional Medical Center",
	"Tulane Medical Center",
	"Centerpoint Medical Center",
	"Rapides Regional Medical Center",
	"Research Psychiatric Center",
	"Lee's Summit Medical Center",
	"John Randolph Medical Center",
	"Spotsylvania Regional Medical Center",
	"CJW Medical Center",
	"LewisGale Hospital Pulaski",
	"Portsmouth Regional Hospital",
	"Reston Hospital Center",
	"Henrico Doctors Hospital",
	"Frankfort Regional Medical Center",
	"LewisGale Medical Center",
	"St. David's Heart Hospital of Austin",
	"St. David's South Austin Medical Center",
	"St. David's Surgical Hospital",
	"St. David's Round Rock Medical Center",
	"St. David's Medical Center",
	"St. David's North Austin Medical Center",
	"Del Sol Medical Center",
	"Las Palmas Medical Center",
	"St. David's Georgetown Hospital",
	"Plantation General Hospital",
	"St. Lucie Medical Center",
	"Westside Regional Medical Center",
	"Aventura Hospital & Medical Center",
	"Northwest Medical Center",
	"Raulerson Hospital - Okeechobee",
	"Palms West Hospital",
	"Mercy Hospital",
	"Lawnwood Regional Medical Center",
	"Kendall Regional Medical Center",
	"JFK Medical Center - Palm Beach",
	"JFK Medical Center North Campus",
	"MountainView Hospital",
	"Riverside Community Hospital",
	"Regional Medical Center of San Jose",
	"Good Samaritan Hospital",
	"Sunrise Hospital",
	"southern hills hospital",
	"West Hills Hospital",
	"Los Robles Regional Medical Center",
	"Thousand Oaks Surgical Hospital",
	"Sky Ridge Medical Center",
	"North Suburban Medical Center",
	"The Medical Center of Aurora",
	"Presbyterian/St. Luke's Medical Center",
	"Surgery Center of the Rockies",
	"Swedish Medical Center",
	"Rose Medical Center",
	"Denver Heart - RMC",
	"Rose Surgical Center",
	"Denver Midtown Surgery Center",
	"HCA Houston Healthcare Southeast",
	"HCA Houston Healthcare Tomball",
	"Rio Grande Regional Hospital",
	"HCA Houston Healthcare Conroe",
	"HCA Houston Healthcare West",
	"HCA Houston Healthcare Mainland",
	"Corpus Christi Medical Center Bay Area",
	"HCA Houston Healthcare Pearland",
	"HCA Houston Healthcare Clear Lake",
	"HCA Houston Healthcare Northwest",
	"Valley Regional Medical Center",
	"HCA Houston Healthcare Kingwood",
	"HCA Houston Healthcare Medical Center",
	"The Woman's Hospital of Texas- Houston",
	"Texas Orthopedic Hospital- Houston",
	"Methodist Specialty & Transplant Hospital",
	"Methodist Hospital",
	"Methodist Hospital Stone Oak",
	"Methodist Hospital Metropolitan",
	"Methodist Hospital South",
	"Methodist Children's Hospital",
	"Methodist Hospital Northeast",
	"Methodist Healthcare System",
	"St. Mark's Hospital",
	"Mountain View Hospital",
	"Alaska Regional Hospital",
	"Eastern Idaho Regional Medical Center",
	"Ogden Regional Medical Center",
	"Timpanogos Regional Hospital",
	"West Valley Medical Group",
	"Osceola Regional Medical Center",
	"Fort Walton Beach Medical Center",
	"North Florida Regional Medical Center",
	"Capital Regional Medical Center",
	"Ocala Regional Medical Center",
	"Poinciana Medical Center",
	"West Florida Hospital",
	"Gulf Coast Regional Medical Center",
	"Central Florida Regional Hospital",
	"West Marion Community Hospital",
	"Putnam Community Medical Center",
	"Oviedo Medical Center",
	"Medical City Dallas",
	"Medical City North Hills",
	"MEDICAL CITY FORT WORTH",
	"Medical City Denton",
	"Medical City Lewisville",
	"Medical City Arlington",
	"Medical City Frisco",
	"Medical City McKinney",
	"Medical City Las Colinas",
	"Medical City Plano",
	"Redmond Regional Medical Center",
	"Parkridge East Hospital",
	"TriStar StoneCrest Medical Center",
	"TriStar Centennial Medical Center",
	"Eastside Medical Center - Snellville",
	"TriStar Southern Hills Medical Center",
	"TriStar Skyline Medical Center",
	"Cartersville Medical Center",
	"TriStar Horizon Medical Center",
	"TriStar Greenview REGIONAL HOSPITAL",
	"TriStar Hendersonville Medical Center",
	"Blake Medical Center",
	"Brandon Regional Hospital",
	"Doctors Hospital of Sarasota",
	"Medical Center of Trinity",
	"Memorial Hospital of Tampa",
	"Largo Medical Center",
	"Regional Medical Center Bayonet Point",
	"Citrus Memorial Hospital",
	"Fawcett Memorial Hospital",
	"St. Petersburg General Hospital",
	"Oak Hill Hospital - Spring Hill",
	"South Bay Hospital - Sun City Center",
	"Largo Medical Center - Indian Rocks",
	"Englewood Community Hospital",
	"Aspen Valley Hospital",
	"Walmart",
	"Sam's Club",
	"DISH",
	"ASSURANCE Independent Agents",
	"Assurance",
	"Veyo",
	"Grubhub",
	"APHRIA",
	"SYKES",
	"Uber",
	"Amazon Workforce Staffing",
	"Uber Eats",
	"Amazon",
	"FedEx",
	"Instacart",
	"Doordash",
	"Uber Eats Headquarters",
	"DoorDash Headquarters",
	"DoorDash",
	"Amazon Headquarters",
	"Shipt Headquarters",
	"Shipt",
	"Lyft"}
