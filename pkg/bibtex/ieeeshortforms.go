package bibtex

var (
	ieeeShortforms = &map[string]string{" Abstracts ": " Abstr. ",
		" Analysis ":               " Anal. ",
		" Academy ":                " Acad. ",
		" Annals ":                 " Ann. ",
		" Accelerator ":            " Accel. ",
		" Annual ":                 " Annu. ",
		" Acoustics ":              " Acoust. ",
		" Apparatus ":              " App. ",
		" Active ":                 " Act. ",
		" Applications ":           " Appl. ",
		" Administration ":         " Admin. ",
		" Applied ":                " Appl. ",
		" Administrative ":         " Administ. ",
		" Approximate ":            " Approx. ",
		" Advanced ":               " Adv. ",
		" Archive ":                " Arch. ",
		" Archives ":               " Arch. ",
		" Aeronautics ":            " Aeronaut. ",
		" Artificial ":             " Artif. ",
		" Aerospace ":              " Aerosp. ",
		" Assembly ":               " Assem. ",
		" Affective ":              " Affect. ",
		" Association ":            " Assoc. ",
		" Africa ":                 " Afr. ",
		" African ":                " Afr. ",
		" Astronomy ":              " Astron. ",
		" Aircraft ":               " Aircr. ",
		" Astronautics ":           " Astronaut. ",
		" Algebraic ":              " Algebr. ",
		" Astrophysics ":           " Astrophys. ",
		" American ":               " Amer. ",
		" Atmosphere ":             " Atmos. ",
		" Atomic ":                 " At. ",
		" Atoms ":                  " At. ",
		" Broadcasting ":           " Broadcast. ",
		" Australasian ":           " Australas. ",
		" Bulletin ":               " Bull. ",
		" Australia ":              " Aust. ",
		" Bureau ":                 " Bur. ",
		" Automatic ":              " Autom. ",
		" Business ":               " Bus. ",
		" Automation ":             " Automat. ",
		" Canadian ":               " Can. ",
		" Automotive ":             " Automot. ",
		" Ceramic ":                " Ceram. ",
		" Autonomous ":             " Auton. ",
		" Chemical ":               " Chem. ",
		" Behavior ":               " Behav. ",
		" Behavioral ":             " Behav. ",
		" Chinese ":                " Chin. ",
		" Belgian ":                " Belg. ",
		" Climatology ":            " Climatol. ",
		" Biochemical ":            " Biochem. ",
		" Clinical ":               " Clin. ",
		" Bioinformatics ":         " Bioinf. ",
		" Cognitive ":              " Cogn. ",
		" Biology ":                " Biol. ",
		" Biological ":             " Biol. ",
		" Colloquium ":             " Colloq. ",
		" Biomedical ":             " Biomed. ",
		" Communications ":         " Commun. ",
		" Biophysics ":             " Biophys. ",
		" Compatibility ":          " Compat. ",
		" British ":                " Brit. ",
		" Component ":              " Compon. ",
		" Components ":             " Compon. ",
		" Computational ":          " Comput. ",
		" Delivery ":               " Del. ",
		" Computer ":               " Comput. ",
		" Computers ":              " Comput. ",
		" Department ":             " Dept. ",
		" Computing ":              " Comput. ",
		" Design ":                 " Des. ",
		" Condensed ":              " Condens. ",
		" Detector ":               " Detect. ",
		" Conference ":             " Conf. ",
		" Development ":            " Develop. ",
		" Congress ":               " Congr. ",
		" Differential ":           " Differ. ",
		" Consumer ":               " Consum. ",
		" Digest ":                 " Dig. ",
		" Conversion ":             " Convers. ",
		" Digital ":                " Digit. ",
		" Convention ":             " Conv. ",
		" Disclosure ":             " Discl. ",
		" Correspondence ":         " Corresp. ",
		" Discussions ":            " Discuss. ",
		" Critical ":               " Crit. ",
		" Dissertations ":          " Diss. ",
		" Crystal ":                " Cryst. ",
		" Distributed ":            " Distrib. ",
		" Crystallography ":        " Crystallogr. ",
		" Dynamics ":               " Dyn. ",
		" Cybernetics ":            " Cybern. ",
		" Earthquake ":             " Earthq. ",
		" Decision ":               " Decis. ",
		" Economic ":               " Econ. ",
		" Economics ":              " Econ. ",
		" Edition ":                " Ed. ",
		" Evolutionary ":           " Evol. ",
		" Education ":              " Educ. ",
		" Exhibition ":             " Exhib. ",
		" Electrical ":             " Elect. ",
		" Experimental ":           " Exp. ",
		" Electrification ":        " Electrific. ",
		" Exploratory ":            " Explor. ",
		" Electromagnetic ":        " Electromagn. ",
		" Exposition ":             " Expo. ",
		" Electroacoustic ":        " Electroacoust. ",
		" Express ":                " Express. ",
		" Electronic ":             " Electron. ",
		" Fabrication ":            " Fabr. ",
		" Emerging ":               " Emerg. ",
		" Faculty ":                " Fac. ",
		" Engineering ":            " Eng. ",
		" Ferroelectrics ":         " Ferroelect. ",
		" Environment ":            " Environ. ",
		" Francais ":               " Fr. ",
		" French ":                 " Fr. ",
		" Equations ":              " Equ. ",
		" Frequency ":              " Freq. ",
		" Equipment ":              " Equip. ",
		" Foundation ":             " Found. ",
		" Ergonomics ":             " Ergonom. ",
		" Fundamental ":            " Fundam. ",
		" European ":               " Eur. ",
		" Generation ":             " Gener. ",
		" Evaluation ":             " Eval. ",
		" Geology ":                " Geol. ",
		" Geophysics ":             " Geophys. ",
		" Innovation ":             " Innov. ",
		" Geoscience ":             " Geosci. ",
		" Institute ":              " Inst. ",
		" Graphics ":               " Graph. ",
		" Instrument ":             " Instrum. ",
		" Guidance ":               " Guid. ",
		" Instrumentation ":        " Instrum. ",
		" Harmonic ":               " Harmon. ",
		" Harmonics ":              " Harmon. ",
		" Insulation ":             " Insul. ",
		" History ":                " Hist. ",
		" Integrated ":             " Integr. ",
		" Horizon ":                " Horiz. ",
		" Intelligence ":           " Intell. ",
		" Hungary ":                " Hung. ",
		" Hungarian ":              " Hung. ",
		" Intelligent ":            " Intell. ",
		" Hydraulics ":             " Hydraul. ",
		" Interactions ":           " Interact. ",
		" Hydrology ":              " Hydrol. ",
		" International ":          " Int. ",
		" Illuminating ":           " Illum. ",
		" Isotopes ":               " Isot. ",
		" Imaging ":                " Imag. ",
		" Israel ":                 " Isr. ",
		" Industrial ":             " Ind. ",
		" Japan ":                  " Jpn. ",
		" Information ":            " Inf. ",
		" Journal ":                " J. ",
		" Informatics ":            " Inform. ",
		" Knowledge ":              " Knowl. ",
		" Laboratory ":             " Lab. ",
		" Laboratoryies ":          " Lab. ",
		" Mathematical ":           " Math. ",
		" Language ":               " Lang. ",
		" Mathematics ":            " Math. ",
		" Learning ":               " Learn. ",
		" Measurement ":            " Meas. ",
		" Letter ":                 " Lett. ",
		" Letters ":                " Lett. ",
		" Mechanical ":             " Mech. ",
		" Lightwave ":              " Lightw. ",
		" Medical ":                " Med. ",
		" Logic ":                  " Log. ",
		" Logical ":                " Log. ",
		" Metals ":                 " Met. ",
		" Luminescence ":           " Lumin. ",
		" Metallurgy ":             " Metall. ",
		" Machine ":                " Mach. ",
		" Meteorology ":            " Meteorol. ",
		" Magazine ":               " Mag. ",
		" Metropolitan ":           " Metrop. ",
		" Magnetics ":              " Magn. ",
		" Mexican ":                " Mex. ",
		" Mexico ":                 " Mex. ",
		" Management ":             " Manage. ",
		" Microelectromechanical ": " Microelectromech. ",
		" Managing ":               " Manag. ",
		" Microgravity ":           " Microgr. ",
		" Manufacturing ":          " Manuf. ",
		" Microscopy ":             " Microsc. ",
		" Marine ":                 " Mar. ",
		" Microwave ":              " Microw. ",
		" Microwaves ":             " Microw. ",
		" Material ":               " Mater. ",
		" Military ":               " Mil. ",
		" Modeling ":               " Model. ",
		" Oceanic ":                " Ocean. ",
		" Molecular ":              " Mol. ",
		" Oceanography ":           " Oceanogr. ",
		" Monitoring ":             " Monit. ",
		" Occupation ":             " Occupat. ",
		" Multiphysics ":           " Multiphys. ",
		" Operational ":            " Oper. ",
		" Nanobioscience ":         " Nanobiosci. ",
		" Optical ":                " Opt. ",
		" Nanotechnology ":         " Nanotechnol. ",
		" Optics ":                 " Opt. ",
		" National ":               " Nat. ",
		" Optimization ":           " Optim. ",
		" Naval ":                  " Nav. ",
		" Organization ":           " Org. ",
		" Network ":                " Netw. ",
		" Networking ":             " Netw. ",
		" Packaging ":              " Packag. ",
		" Newsletter ":             " Newslett. ",
		" Particle ":               " Part. ",
		" Nondestructive ":         " Nondestruct. ",
		" Patent ":                 " Pat. ",
		" Nuclear ":                " Nucl. ",
		" Performance ":            " Perform. ",
		" Numerical ":              " Numer. ",
		" Personal ":               " Pers. ",
		" Observations ":           " Observ. ",
		" Philosophical ":          " Philos. ",
		" Photonics ":              " Photon. ",
		" Productivity ":           " Productiv. ",
		" Photovoltaics ":          " Photovolt. ",
		" Programming ":            " Program. ",
		" Physics ":                " Phys. ",
		" Progress ":               " Prog. ",
		" Physiology ":             " Physiol. ",
		" Propagation ":            " Propag. ",
		" Planetary ":              " Planet. ",
		" Psychology ":             " Psychol. ",
		" Pneumatics ":             " Pneum. ",
		" Quality ":                " Qual. ",
		" Pollution ":              " Pollut. ",
		" Quarterly ":              " Quart. ",
		" Polymer ":                " Polym. ",
		" Radiation ":              " Radiat. ",
		" Polytechnic ":            " Polytech. ",
		" Radiology ":              " Radiol. ",
		" Practice ":               " Pract. ",
		" Reactor ":                " React. ",
		" Precision ":              " Precis. ",
		" Receivers ":              " Receiv. ",
		" Principles ":             " Princ. ",
		" Recognition ":            " Recognit. ",
		" Proceedings ":            " Proc. ",
		" Record ":                 " Rec. ",
		" Processing ":             " Process. ",
		" Rehabilitation ":         " Rehabil. ",
		" Production ":             " Prod. ",
		" Reliability ":            " Rel. ",
		" Report ":                 " Rep. ",
		" Semiconductor ":          " Semicond. ",
		" Research ":               " Res. ",
		" Sensing ":                " Sens. ",
		" Resonance ":              " Reson. ",
		" Series ":                 " Ser. ",
		" Resources ":              " Resour. ",
		" Simulation ":             " Simul. ",
		" Review ":                 " Rev. ",
		" Singapore ":              " Singap. ",
		" Robotics ":               " Robot. ",
		" Sistema ":                " Sist. ",
		" Royal ":                  " Roy. ",
		" Society ":                " Soc. ",
		" Safety ":                 " Saf. ",
		" Sociological ":           " Sociol. ",
		" Satellite ":              " Satell. ",
		" Software ":               " Softw. ",
		" Scandinavian ":           " Scand. ",
		" Solar ":                  " Sol. ",
		" Science ":                " Sci. ",
		" Soviet ":                 " Sov. ",
		" Section ":                " Sect. ",
		" Spectroscopy ":           " Spectrosc. ",
		" Security ":               " Secur. ",
		" Spectrum ":               " Spectr. ",
		" Seismology ":             " Seismol. ",
		" Speculations ":           " Specul. ",
		" Selected ":               " Sel. ",
		" Statistics ":             " Statist. ",
		" Structure ":              " Struct. ",
		" Terrestrial ":            " Terr. ",
		" Studies ":                " Stud. ",
		" Theoretical ":            " Theor. ",
		" Superconductivity ":      " Supercond. ",
		" Transactions ":           " Trans. ",
		" Supplement ":             " Suppl. ",
		" Translation ":            " Transl. ",
		" Surface ":                " Surf. ",
		" Transmission ":           " Transmiss. ",
		" Survey ":                 " Surv. ",
		" Transportation ":         " Transp. ",
		" Sustainable ":            " Sustain. ",
		" Tutorials ":              " Tut. ",
		" Symposium ":              " Symp. ",
		" Ultrasonic ":             " Ultrason. ",
		" Systems ":                " Syst. ",
		" University ":             " Univ. ",
		" Technical ":              " Tech. ",
		" Vacuum ":                 " Vac. ",
		" Techniques ":             " Techn. ",
		" Vehicular ":              " Veh. ",
		" Technology ":             " Technol. ",
		" Vibration ":              " Vib. ",
		" Telecommunications ":     " Telecommun. ",
		" Visual ":                 " Vis. ",
		" Television ":             " Telev. ",
		" Welding ":                " Weld. ",
		" Temperature ":            " Temp. ",
		" Working ":                " Work.",
	}
	ieeeTitleShortforms = &map[string]string{
		" of ":              " ",
		" the ":             " ",
		" on ":              " ",
		"Annals ":           "Ann. ",
		"Proceedings ":      "Proc. ",
		"Annual ":           "Annu. ",
		"Record ":           "Rec. ",
		"Colloquium ":       "Colloq. ",
		"Symposium ":        "Symp. ",
		"Conference ":       "Conf. ",
		"Technical Digest ": "Tech. Dig. ",
		"Congress ":         "Congr. ",
		"Technical Paper ":  "Tech. Paper ",
		"Convention ":       "Conv. ",
		"Digest ":           "Dig. ",
		"Exposition ":       "Expo. ",
		"International ":    "Int. ",
		"National ":         "Nat.  ",
		"First ":            "1st ",
		"Second ":           "2nd ",
		"Third ":            "3rd ",
		"Fourth ":           "4th ",
		"Fifth ":            "5th ",
		"Sixth ":            "6th ",
		"Seventh ":          "7th ",
		"Eigth ":            "8th ",
		"Ninth ":            "9th ",
	}
)
